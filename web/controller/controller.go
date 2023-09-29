package controller

import (
	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-cache-server/usecase/media"
)

type Controller struct {
	media media.MediaUsecase
}

func NewController(media media.MediaUsecase) *Controller {
	return &Controller{
		media: media,
	}
}

func (co *Controller) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/media", co.HandleGetApiMedia)
}
