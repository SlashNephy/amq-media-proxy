package web

import (
	"context"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/SlashNephy/amq-cache-server/config"
	"github.com/SlashNephy/amq-cache-server/logging"
	"github.com/SlashNephy/amq-cache-server/web/controller"
	"github.com/SlashNephy/amq-cache-server/web/middleware/logger"
)

type Server struct {
	e      *echo.Echo
	config *config.Config
}

func NewServer(
	config *config.Config,
	controller *controller.Controller,
	loggerMiddleware *logger.Middleware,
) *Server {
	e := echo.New()
	e.HideBanner = true
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Use(
		middleware.RequestID(),
		loggerMiddleware.Process,
		middleware.RateLimiterWithConfig(
			middleware.RateLimiterConfig{
				Store: middleware.NewRateLimiterMemoryStore(10),
				IdentifierExtractor: func(c echo.Context) (string, error) {
					return c.RealIP(), nil
				},
			},
		),
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:     true,
			LogStatus:  true,
			LogLatency: true,
			LogMethod:  true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				logger := logging.FromContext(c.Request().Context())
				logger.InfoContext(c.Request().Context(), "request",
					slog.Int("status", v.Status),
					slog.String("method", v.Method),
					slog.String("uri", v.URI),
					slog.String("remote_ip", c.RealIP()),
					slog.Float64("latency", float64(v.Latency)/float64(time.Second)),
					slog.Any("err", errors.WithStack(v.Error)),
				)
				return nil
			},
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://animemusicquiz.com"},
			AllowMethods: []string{http.MethodGet},
		}),
		middleware.Secure(),
	)

	controller.RegisterRoutes(e)

	return &Server{
		e:      e,
		config: config,
	}
}

func (s *Server) Start() error {
	return s.e.Start(s.config.ServerAddress)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
