package web

import (
	"github.com/google/wire"

	"github.com/SlashNephy/amq-cache-server/web/controller"
	"github.com/SlashNephy/amq-cache-server/web/middleware/logger"
)

var Set = wire.NewSet(
	NewServer,
	controller.NewController,
	logger.NewMiddleware,
)
