package media

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/fs"
)

type MediaService struct {
	config    *config.Config
	fs        fs.FileSystem
	amqClient AMQClient

	// downloadingMap は url をキーとしてダウンロード中であるかを記録する map。値には意味はない
	downloadingMap map[string]struct{}
	// downloadingMapMutex は downloadingMap にアクセスするときに使用する Mutex
	downloadingMapMutex sync.Mutex
}

func NewMediaService(config *config.Config, fs fs.FileSystem, amqClient AMQClient) *MediaService {
	return &MediaService{
		config:         config,
		fs:             fs,
		amqClient:      amqClient,
		downloadingMap: make(map[string]struct{}),
	}
}

func (s *MediaService) IsDownloading(url string) bool {
	s.downloadingMapMutex.Lock()
	defer s.downloadingMapMutex.Unlock()

	_, ok := s.downloadingMap[url]
	return ok
}

func (s *MediaService) getCachePath(mediaURL string) string {
	filename := filepath.Base(mediaURL)
	path := filepath.Join(s.config.CacheDirectory, filename)

	return strings.ReplaceAll(path, string(filepath.Separator), "/")
}

func (s *MediaService) FindCachedMediaPath(mediaURL string) (string, bool) {
	cachePath := s.getCachePath(mediaURL)
	if ok, _ := s.fs.Exists(cachePath); ok {
		return cachePath, true
	}

	return "", false
}

func (s *MediaService) DownloadMedia(ctx context.Context, mediaURL string, writer io.Writer) error {
	s.lockDownloading(mediaURL)
	defer s.unlockDownloading(mediaURL)

	response, err := s.amqClient.FetchMedia(ctx, mediaURL)
	if err != nil {
		return errors.WithStack(err)
	}

	// キャッシュディレクトリを作成
	cachePath := s.getCachePath(mediaURL)
	if err = os.MkdirAll(filepath.Dir(cachePath), os.ModePerm); err != nil {
		return errors.WithStack(err)
	}

	// キャッシュファイルを作成
	cacheFile, err := os.Create(cachePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer cacheFile.Close()

	// cacheFile と writer に同時にコピーする
	multiWriter := io.MultiWriter(cacheFile, writer)
	if _, err = io.Copy(multiWriter, response.Body); err != nil {
		_ = os.Remove(cachePath)
		return errors.WithStack(err)
	}

	return nil
}

func (s *MediaService) lockDownloading(url string) {
	s.downloadingMapMutex.Lock()
	defer s.downloadingMapMutex.Unlock()

	s.downloadingMap[url] = struct{}{}
}

func (s *MediaService) unlockDownloading(url string) {
	s.downloadingMapMutex.Lock()
	defer s.downloadingMapMutex.Unlock()

	delete(s.downloadingMap, url)
}

var _ MediaUsecase = (*MediaService)(nil)
