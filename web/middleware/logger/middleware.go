package logger

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/SlashNephy/amq-media-proxy/logging"
)

type Middleware struct {
	logger *slog.Logger
}

func NewMiddleware(logger *slog.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (m *Middleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := m.logger
		if requestID := c.Response().Header().Get(echo.HeaderXRequestID); requestID != "" {
			logger = logger.With(slog.String("request_id", requestID))
		}

		ctx := logging.WithContext(c.Request().Context(), logger)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
