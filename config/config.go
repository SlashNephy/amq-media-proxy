package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/pkg/errors"
)

type Config struct {
	LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`
	ServerAddress  string `env:"SERVER_ADDRESS" envDefault:":8080"`
	CacheDirectory string `env:"CACHE_DIRECTORY" envDefault:".cache"`
}

// LoadConfig は環境変数から設定値を読み込む。
func LoadConfig() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}
