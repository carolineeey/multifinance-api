package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	Db  *sql.DB
	URL string
}

// NewClient creates a new client.
func NewClient(db *sql.DB, url string) *Client {
	return &Client{
		Db:  db,
		URL: url,
	}
}

// createContextWithTimeout creates a context.Context with timeout set to timeout.
func (c *Client) createContextWithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// writeObject writes an object to the response writer as JSON.
func (c *Client) writeObject(writer http.ResponseWriter, code int, obj interface{}) (err error) {
	encoder := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	return encoder.Encode(obj)
}

// writeErrorObject writes object as response and send in HTTP error code.
func (c *Client) writeErrorObject(writer http.ResponseWriter, err error, code int, message string) {
	errResp := &ErrorResponse{
		Error:   err,
		Message: message,
	}

	_ = c.writeObject(writer, code, errResp)
}
