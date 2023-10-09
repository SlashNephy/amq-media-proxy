package usecase

import (
	"github.com/google/wire"

	"github.com/SlashNephy/amq-media-proxy/usecase/media"
	"github.com/SlashNephy/amq-media-proxy/usecase/validation"
)

var Set = wire.NewSet(
	media.NewService,
	wire.Bind(new(media.Usecase), new(*media.Service)),

	validation.NewService,
	wire.Bind(new(validation.Usecase), new(*validation.Service)),
)
