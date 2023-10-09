package cloudflare_access

import (
	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/domain/entity"
)

const contextKey = "cloudflare_access"

func GetVisitor(c echo.Context) *entity.User {
	if idToken, ok := c.Get(contextKey).(*entity.User); ok {
		return idToken
	}

	return nil
}
