package usecase

import (
	"github.com/google/wire"

	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

var Set = wire.NewSet(
	media.NewMediaService,
	wire.Bind(new(media.MediaUsecase), new(*media.MediaService)),
)
