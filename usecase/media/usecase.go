package media

import (
	"context"
	"io"
	"net/http"
)

type MediaUsecase interface {
	FindCachedMediaPath(ctx context.Context, mediaURL string) (string, bool)
	DownloadMedia(ctx context.Context, mediaURL string, writer io.Writer) error
}

// AMQClient は AMQ ユーザーを装って HTTP リクエストを実行するクライアント
type AMQClient interface {
	FetchMedia(ctx context.Context, mediaURL string) (*http.Response, error)
}
