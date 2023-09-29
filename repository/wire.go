package repository

import (
	"github.com/google/wire"

	"github.com/SlashNephy/amq-media-proxy/repository/external/amq"
	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

var Set = wire.NewSet(
	amq.NewAMQClient,
	wire.Bind(new(media.AMQClient), new(*amq.Client)),
)
