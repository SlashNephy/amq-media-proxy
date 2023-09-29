package usecase

import "github.com/SlashNephy/amq-cache-server/usecase/media"

//go:generate bulkmockgen MockRepos -- -typed -package mock_repo -destination ./mock_repo/mock_repo.go
var MockRepos = []any{
	new(media.AMQClient),
}

//go:generate bulkmockgen MockUsecases -- -typed -package mock_usecase -destination ./mock_usecase/mock_usecase.go
var MockUsecases = []any{
	new(media.MediaUsecase),
}
