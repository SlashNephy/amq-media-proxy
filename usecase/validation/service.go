package validation

import (
	"regexp"

	"github.com/SlashNephy/amq-media-proxy/config"
)

type Service struct {
	referer         string
	mediaURLPattern *regexp.Regexp
}

func NewService(config *config.Config) (*Service, error) {
	mediaURLPattern, err := regexp.Compile(config.MediaURLPattern)
	if err != nil {
		return nil, err
	}

	return &Service{
		referer:         config.ValidReferer,
		mediaURLPattern: mediaURLPattern,
	}, nil
}

func (s Service) CheckReferer(referer string) bool {
	return s.referer == referer
}

func (s Service) CheckMediaURL(url string) bool {
	return s.mediaURLPattern.MatchString(url)
}

var _ Usecase = (*Service)(nil)
