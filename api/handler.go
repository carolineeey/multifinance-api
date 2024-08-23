package api

import (
	"net/http"
	"time"
)

const (
	TimeoutDefault           = 30 * time.Second
	TimeoutHandleGetCustomer = TimeoutDefault
)

// HandleGetCustomer handles getting a customer detail.
func (c *Client) HandleGetCustomer(writer http.ResponseWriter, request *http.Request) {
	_, cancel := c.createContextWithTimeout(TimeoutHandleGetCustomer)
	defer cancel()
}
