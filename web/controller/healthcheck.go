package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (co *Controller) HandleGetHealthcheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
