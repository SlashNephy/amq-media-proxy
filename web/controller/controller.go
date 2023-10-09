package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/usecase/media"
	"github.com/SlashNephy/amq-media-proxy/usecase/validation"
)

type Controller struct {
	media      media.Usecase
	validation validation.Usecase
	config     *config.Config
}

func NewController(
	media media.Usecase,
	validation validation.Usecase,
	config *config.Config,
) *Controller {
	return &Controller{
		media:      media,
		validation: validation,
		config:     config,
	}
}

func (co *Controller) RegisterRoutes(e *echo.Echo) {
	e.GET("/", co.HandleGetIndex)

	e.GET("/healthcheck", co.HandleGetHealthcheck)

	api := e.Group("/api")
	api.GET("/media", co.HandleGetApiMedia)
}
