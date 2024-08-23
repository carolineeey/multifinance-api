package api

type Client struct {
	URL string
}

// NewClient creates a new client with a default logger.
func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}
