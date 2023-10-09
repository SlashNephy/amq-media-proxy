package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/SlashNephy/amq-media-proxy/logging"
	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

type Downloader struct {
	media media.MediaUsecase
	eg    *errgroup.Group
	ctx   context.Context
}

func NewDownloader(ctx context.Context, media media.MediaUsecase, limit int) *Downloader {
	eg, egctx := errgroup.WithContext(ctx)
	eg.SetLimit(limit)

	return &Downloader{
		media: media,
		eg:    eg,
		ctx:   egctx,
	}
}

func (d *Downloader) QueueDownload(urls []string) {
	total := len(urls)
	for i, url := range urls {
		current := i
		u := url
		d.eg.Go(func() error {
			return d.download(u, current, total)
		})
	}
}

func (d *Downloader) download(url string, current, total int) error {
	// キャンセルされていたら直ちに終了
	if errors.Is(d.ctx.Err(), context.Canceled) {
		return nil
	}

	// キャッシュがあったらスキップ
	if _, ok := d.media.FindCachedMediaPath(url); ok {
		return nil
	}

	if err := d.media.DownloadMedia(context.WithoutCancel(d.ctx), url); err != nil {
		logging.FromContext(d.ctx).Warn("failed to download", slog.String("url", url), slog.Any("err", err))
		return nil
	}

	logging.FromContext(d.ctx).Info("downloaded",
		slog.String("url", url),
		slog.Int("current", current),
		slog.Int("total", total),
	)

	time.Sleep(1 * time.Second)
	return nil
}

func (d *Downloader) Wait() error {
	return d.eg.Wait()
}
