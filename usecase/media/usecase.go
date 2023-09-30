package media

import (
	"context"
	"net/http"
)

type MediaUsecase interface {
	IsValidURL(url string) bool
	IsDownloading(url string) bool
	FindCachedMediaPath(mediaURL string) (string, bool)
	DownloadMedia(ctx context.Context, mediaURL, cachePath string) error
}

// AMQClient は AMQ ユーザーを装って HTTP リクエストを実行するクライアント
type AMQClient interface {
	FetchMedia(ctx context.Context, mediaURL string) (*http.Response, error)
}
