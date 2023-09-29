package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("環境変数名と構造体のフィールドの対応が付いている", func(t *testing.T) {
		tests := []struct {
			Key    string
			Value  string
			Actual func(c *Config) any
		}{
			{
				Key:   "LOG_LEVEL",
				Value: "debug",
				Actual: func(c *Config) any {
					return c.LogLevel
				},
			},
			{
				Key:   "SERVER_ADDRESS",
				Value: ":443",
				Actual: func(c *Config) any {
					return c.ServerAddress
				},
			},
			{
				Key:   "CACHE_DIRECTORY",
				Value: "tmp",
				Actual: func(c *Config) any {
					return c.CacheDirectory
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.Key, func(t *testing.T) {
				err := os.Setenv(tt.Key, tt.Value)
				require.NoError(t, err)

				cfg, err := LoadConfig()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, tt.Actual(cfg))
			})
		}
	})
}
