package repository

import (
	"github.com/google/wire"

	"github.com/SlashNephy/amq-cache-server/repository/external/amq"
	"github.com/SlashNephy/amq-cache-server/usecase/media"
)

var Set = wire.NewSet(
	amq.NewAMQClient,
	wire.Bind(new(media.AMQClient), new(*amq.Client)),
)
