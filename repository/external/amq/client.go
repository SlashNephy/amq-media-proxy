package amq

import "net/http"

type Client struct {
	httpClient *http.Client
}

func NewAMQClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}
