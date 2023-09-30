package controller

import (
	"log/slog"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/domain/content_type"
	"github.com/SlashNephy/amq-media-proxy/logging"
)

var mediaURLPattern = regexp.MustCompile(`^https://\w+\.catbox\.video/\w+\.(?:mp3|webm)$`)

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
	if co.media.IsDownloading(params.URL) {
		return c.Redirect(http.StatusFound, params.URL)
	}

	// HTTP ヘッダーを書き込む
	{
		// MIME Type を判定する
		contentType, err := content_type.DetectContentTypeByFilename(params.URL)
		if err != nil {
			logging.FromContext(c.Request().Context()).Error("unexpected content type",
				slog.String("url", params.URL),
				slog.Any("err", err),
			)
			return echo.ErrBadRequest
		}

		c.Response().Header().Set("Content-Type", string(contentType))
		c.Response().Header().Set("Cache-Control", "public, immutable, max-age=2592000, stale-if-error=604800, stale-while-revalidate=604800")
		c.Response().WriteHeader(http.StatusOK)
	}

	// キャッシュ済みならそれを返す
	if cachePath, ok := co.media.FindCachedMediaPath(params.URL); ok {
		logging.FromContext(c.Request().Context()).Info("found from cache",
			slog.String("url", params.URL),
		)
		return c.File(cachePath)
	}

	// URL をダウンロードしつつ、レスポンスに書き込む
	if err := co.media.DownloadMedia(c.Request().Context(), params.URL, c.Response().Writer); err != nil {
		logging.FromContext(c.Request().Context()).Error("failed to download",
			slog.String("url", params.URL),
			slog.Any("err", err),
		)
		return echo.ErrInternalServerError
	}

	return nil
}
