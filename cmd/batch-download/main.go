package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/fs"
	"github.com/SlashNephy/amq-media-proxy/logging"
	"github.com/SlashNephy/amq-media-proxy/repository/external/amq"
	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, err := logging.NewLogger(cfg)
	if err != nil {
		panic(err)
	}

	urls, err := LoadMediaURLs(cfg.QuestionsJSONPath)
	if err != nil {
		panic(err)
	}

	media, err := media.NewService(cfg, fs.NewRealFileSystem(), amq.NewAMQClient())
	if err != nil {
		panic(err)
	}

	downloader := NewDownloader(logging.WithContext(ctx, logger), media, 3)
	downloader.QueueDownload(urls)
	if err = downloader.Wait(); err != nil {
		panic(err)
	}
}
