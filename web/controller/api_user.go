package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	cloudflareAccess "github.com/SlashNephy/amq-media-proxy/web/middleware/cloudflare_access"
)

func (co *Controller) handleGetApiUser(c echo.Context) error {
	visitor := cloudflareAccess.GetVisitor(c)
	if visitor == nil {
		return c.JSON(http.StatusUnauthorized, map[string]bool{
			"success": false,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success":  true,
		"id":       visitor.ID,
		"username": visitor.Username,
	})
}
