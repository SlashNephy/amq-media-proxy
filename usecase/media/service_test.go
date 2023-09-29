package media

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/SlashNephy/amq-cache-server/config"
	"github.com/SlashNephy/amq-cache-server/fs"
	"github.com/SlashNephy/amq-cache-server/usecase/mock_repo"
)

type MediaServiceMocks struct {
	fs        fs.FileSystem
	amqClient *mock_repo.MockAMQClient
}

func NewTestMediaService(t *testing.T, cfg *config.Config, exists bool) (*MediaService, *MediaServiceMocks) {
	ctrl := gomock.NewController(t)
	m := &MediaServiceMocks{
		amqClient: mock_repo.NewMockAMQClient(ctrl),
	}
	s := NewMediaService(cfg, fs.NewFakeFileSystem(exists), m.amqClient)
	return s, m
}

func TestMediaService_FindCachedMediaPath(t *testing.T) {
	t.Run("キャッシュが存在しているときに true が返る", func(t *testing.T) {
		s, _ := NewTestMediaService(t, &config.Config{
			CacheDirectory: "/tmp",
		}, true)

		path, ok := s.FindCachedMediaPath(context.Background(), "https://catbox.video/challenge.mp3")
		require.Equal(t, path, "/tmp/challenge.mp3")
		assert.True(t, ok)
	})

	t.Run("キャッシュが存在していないときに false が返る", func(t *testing.T) {
		s, _ := NewTestMediaService(t, &config.Config{
			CacheDirectory: "/tmp",
		}, false)

		path, ok := s.FindCachedMediaPath(context.Background(), "https://catbox.video/challenge.mp3")
		require.Equal(t, path, "")
		assert.False(t, ok)
	})
}
