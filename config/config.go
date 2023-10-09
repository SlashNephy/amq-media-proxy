package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/pkg/errors"
)

type Config struct {
	LogLevel       string `env:"LOG_LEVEL" envDefault:"info"`
	ServerAddress  string `env:"SERVER_ADDRESS" envDefault:":8080"`
	CacheDirectory string `env:"CACHE_DIRECTORY" envDefault:".cache"`
	TrustRealIP    bool   `env:"TRUST_REAL_IP"`

	MediaURLPattern string `env:"MEDIA_URL_PATTERN"`
	ValidReferer    string `env:"VALID_REFERER"`
	ValidOrigin     string `env:"VALID_ORIGIN"`

	QuestionsJSONPath string `env:"QUESTIONS_JSON_PATH"`
}

// LoadConfig は環境変数から設定値を読み込む。
func LoadConfig() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}
