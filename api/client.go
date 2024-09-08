package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/schema"
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

// decodeRequestJson decodes HTTP request body as JSON to `obj` using json pkg.
func (c *Client) decodeRequestJson(writer http.ResponseWriter, request *http.Request, obj interface{}) (err error) {
	if request.Body == nil {
		c.writeErrorObject(writer, err, http.StatusBadRequest, "request body is nil")
		return
	}

	err = json.NewDecoder(request.Body).Decode(obj)
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to decode request body")
		return
	}

	return nil
}

// decodeRequestSchema decodes HTTP request query string to `obj` using gorilla/schema pkg.
func (c *Client) decodeRequestSchema(writer http.ResponseWriter, request *http.Request, obj interface{}) (err error) {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err = decoder.Decode(obj, request.URL.Query())
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to parse request query parameter")
		return err
	}

	return nil
}
