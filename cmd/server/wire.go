//go:build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"github.com/SlashNephy/amq-cache-server/config"
	"github.com/SlashNephy/amq-cache-server/fs"
	"github.com/SlashNephy/amq-cache-server/logging"
	"github.com/SlashNephy/amq-cache-server/repository"
	"github.com/SlashNephy/amq-cache-server/usecase"
	"github.com/SlashNephy/amq-cache-server/web"
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
