package media

import (
	"context"
	"github.com/SlashNephy/amq-cache-server/fs"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/SlashNephy/amq-cache-server/config"
	"github.com/SlashNephy/amq-cache-server/logging"
)

type MediaService struct {
	config    *config.Config
	fs        fs.FileSystem
	amqClient AMQClient
}

func NewMediaService(config *config.Config, fs fs.FileSystem, amqClient AMQClient) *MediaService {
	return &MediaService{
		config:    config,
		fs:        fs,
		amqClient: amqClient,
	}
}

func (s *MediaService) getCachePath(mediaURL string) string {
	filename := filepath.Base(mediaURL)
	path := filepath.Join(s.config.CacheDirectory, filename)

	return strings.ReplaceAll(path, string(filepath.Separator), "/")
}

func (s *MediaService) FindCachedMediaPath(ctx context.Context, mediaURL string) (string, bool) {
	cachePath := s.getCachePath(mediaURL)
	if ok, _ := s.fs.Exists(cachePath); ok {
		logging.FromContext(ctx).Info("found from cache", slog.String("url", mediaURL))
		return cachePath, true
	}

	return "", false
}

func (s *MediaService) DownloadMedia(ctx context.Context, mediaURL string, writer io.Writer) error {
	response, err := s.amqClient.FetchMedia(ctx, mediaURL)
	if err != nil {
		logging.FromContext(ctx).Info("failed to download url", slog.String("url", mediaURL))
		return err
	}

	// キャッシュファイル
	cachePath := s.getCachePath(mediaURL)
	cacheFile, err := os.Create(cachePath)
	if err != nil {
		return err
	}
	defer cacheFile.Close()

	// cacheFile と writer に同時にコピーする
	// TODO: 同時に同じ URL にリクエストが来たらどうする？
	multiWriter := io.MultiWriter(cacheFile, writer)
	_, err = io.Copy(multiWriter, response.Body)
	return err
}

var _ MediaUsecase = (*MediaService)(nil)
