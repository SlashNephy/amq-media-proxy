package controller

import (
	"log/slog"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-cache-server/domain/content_type"
	"github.com/SlashNephy/amq-cache-server/logging"
)

var mediaURLPattern = regexp.MustCompile(`^https://\w+\.catbox\.video/\w+\.(?:mp3|webm)$`)

func (co *Controller) HandleGetApiMedia(c echo.Context) error {
	var params struct {
		URL string `query:"url"`
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	// 不正な URL が来ないかバリデーション
	if !mediaURLPattern.MatchString(params.URL) {
		logging.FromContext(c.Request().Context()).Error("unexpected url", slog.String("url", params.URL))
		return echo.ErrBadRequest
	}

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
	if err = co.media.DownloadMedia(c.Request().Context(), params.URL, c.Response().Writer); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}
