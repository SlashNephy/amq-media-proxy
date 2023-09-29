//go:build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/fs"
	"github.com/SlashNephy/amq-media-proxy/logging"
	"github.com/SlashNephy/amq-media-proxy/repository"
	"github.com/SlashNephy/amq-media-proxy/usecase"
	"github.com/SlashNephy/amq-media-proxy/web"
)

func InitializeServer(ctx context.Context) (*web.Server, error) {
	wire.Build(
		config.LoadConfig,
		logging.NewLogger,
		web.Set,
		repository.Set,
		usecase.Set,
		fs.Set,
	)

	return nil, nil
}
