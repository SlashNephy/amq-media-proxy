package controller

import (
	"log/slog"
	"net/http"
	"regexp"
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/domain/content_type"
	"github.com/SlashNephy/amq-media-proxy/logging"
)

var (
	mediaURLPattern     = regexp.MustCompile(`^https://\w+\.catbox\.video/\w+\.(?:mp3|webm)$`)
	downloadingMap      = make(map[string]struct{})
	downloadingMapMutex *sync.Mutex
)

func (co *Controller) HandleGetApiMedia(c echo.Context) error {
	var params struct {
		URL string `query:"u"`
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	// 不正な URL が来ないかバリデーション
	if !mediaURLPattern.MatchString(params.URL) {
		logging.FromContext(c.Request().Context()).Error("unexpected url", slog.String("url", params.URL))
		return echo.ErrBadRequest
	}

	// ダウンロード中ならリダイレクトする
	downloadingMapMutex.Lock()
	if _, ok := downloadingMap[params.URL]; ok {
		return c.Redirect(http.StatusFound, params.URL)
	}
	downloadingMap[params.URL] = struct{}{}
	downloadingMapMutex.Unlock()

	// キャッシュ済みならそれを返す
	cachePath, ok := co.media.FindCachedMediaPath(c.Request().Context(), params.URL)
	if ok {
		return c.File(cachePath)
	}

	// MIME Type を判定する
	contentType, err := content_type.DetectContentTypeByFilename(params.URL)
	if err != nil {
		return err
	}

	// HTTP ヘッダーを書き込む
	c.Response().Header().Set("Content-Type", string(contentType))
	c.Response().Header().Set("Cache-Control", "public, immutable, max-age=2592000, stale-if-error=604800, stale-while-revalidate=604800")
	c.Response().WriteHeader(http.StatusOK)

	// URL をダウンロードしつつ、レスポンスに書き込む
	err = co.media.DownloadMedia(c.Request().Context(), params.URL, c.Response().Writer)

	// ダウンロードが終わった
	downloadingMapMutex.Lock()
	delete(downloadingMap, params.URL)
	downloadingMapMutex.Unlock()

	if err != nil {
		return echo.ErrInternalServerError
	}
	return nil
}
