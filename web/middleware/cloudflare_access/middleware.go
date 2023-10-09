package cloudflare_access

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/logging"
)

type Middleware struct {
	verifier *oidc.IDTokenVerifier
}

func NewMiddleware(ctx context.Context, config *config.Config) *Middleware {
	var verifier *oidc.IDTokenVerifier
	if config.CloudflareAccessTeamDomain != "" && config.CloudflareAccessPolicyAudience != "" {
		certsURL := fmt.Sprintf("%s/cdn-cgi/access/certs", config.CloudflareAccessTeamDomain)
		keySet := oidc.NewRemoteKeySet(ctx, certsURL)
		verifier = oidc.NewVerifier(
			config.CloudflareAccessTeamDomain,
			keySet,
			&oidc.Config{
				ClientID: config.CloudflareAccessPolicyAudience,
			},
		)
	}

	return &Middleware{
		verifier: verifier,
	}
}

func (m *Middleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if m.verifier == nil {
			return next(c)
		}

		jwt := getCloudflareAccessJWT(c)
		if jwt == "" {
			return next(c)
		}

		idToken, err := m.verifier.Verify(c.Request().Context(), jwt)
		if err != nil {
			logging.FromContext(c.Request().Context()).Error("failed to verify id_token", slog.Any("err", err))
			return next(c)
		}

		c.Set(contextKey, idToken)
		return next(c)
	}
}

func getCloudflareAccessJWT(c echo.Context) string {
	jwt := c.Request().Header.Get("Cf-Access-Jwt-Assertion")
	if jwt != "" {
		return jwt
	}

	cookie, _ := c.Cookie("CF_Authorization")
	if cookie != nil {
		return cookie.Value
	}

	return ""
}
