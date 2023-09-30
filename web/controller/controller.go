package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

type Controller struct {
	media  media.MediaUsecase
	config *config.Config
}

func NewController(media media.MediaUsecase, config *config.Config) *Controller {
	return &Controller{
		media:  media,
		config: config,
	}
}

func (co *Controller) RegisterRoutes(e *echo.Echo) {
	e.GET("/healthcheck", co.HandleGetHealthcheck)

	api := e.Group("/api")
	api.GET("/media", co.HandleGetApiMedia)
}
