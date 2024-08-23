package api

import _ "github.com/carolineeey/multifinance-api/api"

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}
