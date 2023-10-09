package controller

import (
	cloudflareAccess "github.com/SlashNephy/amq-media-proxy/web/middleware/cloudflare_access"
	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/usecase/media"
	"github.com/SlashNephy/amq-media-proxy/usecase/validation"
)

type Controller struct {
	media                      media.Usecase
	validation                 validation.Usecase
	cloudflareAccessMiddleware *cloudflareAccess.Middleware
}

func NewController(
	media media.Usecase,
	validation validation.Usecase,
	cloudflareAccessMiddleware *cloudflareAccess.Middleware,
) *Controller {
	return &Controller{
		media:                      media,
		validation:                 validation,
		cloudflareAccessMiddleware: cloudflareAccessMiddleware,
	}
}

func (co *Controller) RegisterRoutes(e *echo.Echo) {
	e.GET("/healthcheck", co.HandleGetHealthcheck)

	api := e.Group("/api", co.cloudflareAccessMiddleware.Process)
	api.GET("/media", co.HandleGetApiMedia)
}
