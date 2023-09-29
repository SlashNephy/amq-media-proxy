package amq

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	"github.com/SlashNephy/amq-media-proxy/usecase/media"
)

var requestHeaders = map[string]string{
	"Accept":             "*/*",
	"DNT":                "1",
	"Referer":            "https://animemusicquiz.com/",
	"Sec-Fetch-Dest":     "video",
	"Sec-Fetch-Mode":     "no-cors",
	"Sec-Fetch-Site":     "cross-site",
	"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	"sec-ch-ua":          `"Google Chrome";v="117", "Not;A=Brand";v="8", "Chromium";v="117"`,
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": `"Windows"`,
}

func (c *Client) FetchMedia(ctx context.Context, mediaURL string) (*http.Response, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, mediaURL, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for k, v := range requestHeaders {
		request.Header.Set(k, v)
	}

	return c.httpClient.Do(request)
}

var _ media.AMQClient = (*Client)(nil)
