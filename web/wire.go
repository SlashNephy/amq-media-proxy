package web

import (
	"github.com/google/wire"

	"github.com/SlashNephy/amq-media-proxy/web/controller"
	"github.com/SlashNephy/amq-media-proxy/web/middleware/logger"
)

var Set = wire.NewSet(
	NewServer,
	controller.NewController,
	logger.NewMiddleware,
)
