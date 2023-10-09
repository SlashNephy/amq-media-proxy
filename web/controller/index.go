package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (co *Controller) HandleGetIndex(c echo.Context) error {
	return c.Redirect(http.StatusFound, co.config.ValidOrigin)
}
