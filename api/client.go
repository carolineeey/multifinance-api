package api

import (
	"context"
	"time"
)

type Client struct {
	URL string
}

// NewClient creates a new client with a default logger.
func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

// createContextWithTimeout creates a context.Context with timeout set to timeout.
func (c *Client) createContextWithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
