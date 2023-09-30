package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"

	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

type Downloader struct {
	media media.MediaUsecase
	eg    *errgroup.Group
	ctx   context.Context
	bar   *progressbar.ProgressBar
}

func NewDownloader(ctx context.Context, media media.MediaUsecase, limit int) *Downloader {
	eg, egctx := errgroup.WithContext(ctx)
	eg.SetLimit(limit)

	return &Downloader{
		media: media,
		eg:    eg,
		ctx:   egctx,
		bar:   progressbar.Default(-1),
	}
}

func (d *Downloader) QueueDownload(urls []string) {
	d.bar.ChangeMax(len(urls))

	for _, url := range urls {
		u := url
		d.eg.Go(func() error {
			return d.download(u)
		})
	}
}

func (d *Downloader) download(url string) error {
	// キャンセルされていたら直ちに終了
	if errors.Is(d.ctx.Err(), context.Canceled) {
		return nil
	}

	if err := d.bar.Add(1); err != nil {
		return err
	}

	// キャッシュがあったらスキップ
	if _, ok := d.media.FindCachedMediaPath(url); ok {
		return nil
	}

	if err := d.media.DownloadMedia(context.WithoutCancel(d.ctx), url); err != nil {
		return fmt.Errorf("failed to download %s: %w", url, errors.WithStack(err))
	}

	time.Sleep(1 * time.Second)
	return nil
}

func (d *Downloader) Wait() error {
	return d.eg.Wait()
}
