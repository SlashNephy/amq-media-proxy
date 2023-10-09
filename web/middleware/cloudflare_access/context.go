package cloudflare_access

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
)

const contextKey = "cloudflare_access"

func GetVisitor(c echo.Context) *oidc.IDToken {
	if idToken, ok := c.Get(contextKey).(*oidc.IDToken); ok {
		return idToken
	}

	return nil
}
