package logging

import (
	"log/slog"
	"os"

	"github.com/pkg/errors"

	"github.com/SlashNephy/amq-cache-server/config"
)

func NewLogger(cfg *config.Config) (*slog.Logger, error) {
	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(cfg.LogLevel)); err != nil {
		return nil, errors.WithStack(err)
	}

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: &logLevel,
	})), nil
}
