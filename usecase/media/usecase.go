package media

import (
	"context"
	"net/http"
)

type Usecase interface {
	FindCachedMediaPath(mediaURL string) (string, bool)
	DownloadMedia(ctx context.Context, mediaURL string) error
}

// AMQClient は AMQ ユーザーを装って HTTP リクエストを実行するクライアント
type AMQClient interface {
	FetchMedia(ctx context.Context, mediaURL string) (*http.Response, error)
}
