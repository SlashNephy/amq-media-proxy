// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/fs"
	"github.com/SlashNephy/amq-media-proxy/logging"
	"github.com/SlashNephy/amq-media-proxy/repository/external/amq"
	"github.com/SlashNephy/amq-media-proxy/usecase/media"
	"github.com/SlashNephy/amq-media-proxy/usecase/validation"
	"github.com/SlashNephy/amq-media-proxy/web"
	"github.com/SlashNephy/amq-media-proxy/web/controller"
	"github.com/SlashNephy/amq-media-proxy/web/middleware/logger"
)

// Injectors from wire.go:

func InitializeServer(ctx context.Context) (*web.Server, error) {
	configConfig, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	realFileSystem := fs.NewRealFileSystem()
	client := amq.NewAMQClient()
	service, err := media.NewService(configConfig, realFileSystem, client)
	if err != nil {
		return nil, err
	}
	validationService, err := validation.NewService(configConfig)
	if err != nil {
		return nil, err
	}
	controllerController := controller.NewController(service, validationService)
	slogLogger, err := logging.NewLogger(configConfig)
	if err != nil {
		return nil, err
	}
	middleware := logger.NewMiddleware(slogLogger)
	server := web.NewServer(configConfig, controllerController, middleware)
	return server, nil
}
