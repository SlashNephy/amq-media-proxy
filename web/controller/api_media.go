package controller

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/domain/content_type"
	"github.com/SlashNephy/amq-media-proxy/logging"
)

func (co *Controller) HandleGetApiMedia(c echo.Context) error {
	var params struct {
		URL string `query:"u"`
	}
	if err := c.Bind(&params); err != nil {
		return err
	}

	// 不正な URL が来ないかバリデーション
	if !co.media.IsValidURL(params.URL) {
		logging.FromContext(c.Request().Context()).Error("unexpected url", slog.String("url", params.URL))
		return echo.ErrBadRequest
	}

	// キャッシュ済みならファイルを送信する
	if cachePath, ok := co.media.FindCachedMediaPath(params.URL); ok {
		logging.FromContext(c.Request().Context()).Info("found from cache", slog.String("url", params.URL))

		// MIME Type を判定する
		contentType, err := content_type.DetectContentTypeByFilename(params.URL)
		if err != nil {
			logging.FromContext(c.Request().Context()).Error("unexpected content type",
				slog.String("url", params.URL),
				slog.Any("err", err),
			)
			return echo.ErrBadRequest
		}

		// HTTP ヘッダーを書き込む
		c.Response().Header().Set("Content-Type", string(contentType))
		c.Response().Header().Set("Cache-Control", "public, immutable, max-age=2592000, stale-if-error=604800, stale-while-revalidate=604800")
		c.Response().WriteHeader(http.StatusOK)
		return c.File(cachePath)
	}

	// ダウンロード
	go func(ctx context.Context, url string) {
		if err := co.media.DownloadMedia(ctx, url); err != nil {
			logging.FromContext(ctx).Error("failed to download",
				slog.String("url", url),
				slog.Any("err", err),
			)
		} else {
			logging.FromContext(ctx).Info("downloaded", slog.String("url", url))
		}
	}(context.WithoutCancel(c.Request().Context()), params.URL)

	// リダイレクト
	return c.Redirect(http.StatusFound, params.URL)
}
