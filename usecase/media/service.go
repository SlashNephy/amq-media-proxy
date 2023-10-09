package media

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pkg/errors"

	"github.com/SlashNephy/amq-media-proxy/config"
	"github.com/SlashNephy/amq-media-proxy/fs"
)

type Service struct {
	config    *config.Config
	fs        fs.FileSystem
	amqClient AMQClient

	// downloadingMap は url をキーとしてダウンロード中であるかを記録する map。値には意味はない
	downloadingMap map[string]struct{}
	// downloadingMapMutex は downloadingMap にアクセスするときに使用する Mutex
	downloadingMapMutex sync.Mutex
}

func NewService(config *config.Config, fs fs.FileSystem, amqClient AMQClient) (*Service, error) {
	return &Service{
		config:         config,
		fs:             fs,
		amqClient:      amqClient,
		downloadingMap: make(map[string]struct{}),
	}, nil
}

func (s *Service) getCachePath(mediaURL string) string {
	// キャッシュディレクトリを作成
	if err := os.MkdirAll(s.config.CacheDirectory, os.ModePerm); err != nil {
		panic(err)
	}

	filename := filepath.Base(mediaURL)
	path := filepath.Join(s.config.CacheDirectory, filename)
	path = strings.ReplaceAll(path, string(filepath.Separator), "/")
	return path
}

func (s *Service) FindCachedMediaPath(mediaURL string) (string, bool) {
	cachePath := s.getCachePath(mediaURL)
	if ok, _ := s.fs.Exists(cachePath); ok {
		return cachePath, true
	}

	return "", false
}

func (s *Service) DownloadMedia(ctx context.Context, mediaURL string) error {
	// すでにダウンロード中なら何もしない
	if s.isDownloading(mediaURL) {
		return nil
	}

	s.lockDownloading(mediaURL)
	defer s.unlockDownloading(mediaURL)

	response, err := s.amqClient.FetchMedia(ctx, mediaURL)
	if err != nil {
		return errors.WithStack(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid response, %d", response.StatusCode)
	}

	// tmpCachePath が存在していたら削除
	cachePath := s.getCachePath(mediaURL)
	tmpCachePath := cachePath + ".tmp"
	if ok, _ := s.fs.Exists(tmpCachePath); ok {
		if err = os.Remove(tmpCachePath); err != nil {
			return errors.WithStack(err)
		}
	}

	// キャッシュファイルを作成
	cacheFile, err := os.Create(tmpCachePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer os.Remove(tmpCachePath)

	if _, err = io.Copy(cacheFile, response.Body); err != nil {
		return errors.WithStack(err)
	}

	if err = cacheFile.Close(); err != nil {
		return errors.WithStack(err)
	}

	return os.Rename(tmpCachePath, cachePath)
}

func (s *Service) isDownloading(url string) bool {
	s.downloadingMapMutex.Lock()
	defer s.downloadingMapMutex.Unlock()

	_, ok := s.downloadingMap[url]
	return ok
}

func (s *Service) lockDownloading(url string) {
	s.downloadingMapMutex.Lock()
	defer s.downloadingMapMutex.Unlock()

	s.downloadingMap[url] = struct{}{}
}

func (s *Service) unlockDownloading(url string) {
	s.downloadingMapMutex.Lock()
	defer s.downloadingMapMutex.Unlock()

	delete(s.downloadingMap, url)
}

var _ Usecase = (*Service)(nil)
