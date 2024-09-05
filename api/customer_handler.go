package api

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	TimeoutDefault              = 30 * time.Second
	TimeoutHandleGetCustomer    = TimeoutDefault
	TimeoutHandleCreateCustomer = TimeoutDefault
	TimeoutHandleEditCustomer   = TimeoutDefault
	TimeoutHandleDeleteCustomer = TimeoutDefault
)

// HandleGetCustomer handles getting a customer detail.
func (c *Client) HandleGetCustomer(writer http.ResponseWriter, request *http.Request) {
	_, cancel := c.createContextWithTimeout(TimeoutHandleGetCustomer)
	defer cancel()
}

// HandleAddCustomer handles creating a new customer.
func (c *Client) HandleAddCustomer(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := c.createContextWithTimeout(TimeoutHandleCreateCustomer)
	defer cancel()

	payload := &SetUserRequest{}

	if err := json.NewDecoder(request.Body).Decode(payload); err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to decode payload")
		return
	}

	err := c.CreateUserCtx(ctx, payload)
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to create user")
		return
	}

	response := map[string]int{
		"nik": payload.Nik,
	}

	_ = c.writeObject(writer, http.StatusOK, response)
}

// HandleEditCustomer handles request to edit the data of a customer.
func (c *Client) HandleEditCustomer(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := c.createContextWithTimeout(TimeoutHandleEditCustomer)
	defer cancel()

	payload := &SetUserRequest{}

	if err := json.NewDecoder(request.Body).Decode(payload); err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to decode payload")
		return
	}

	err := c.EditUserDataCtx(ctx, payload)
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to edit user data")
		return
	}

	response := map[string]int{
		"nik": payload.Nik,
	}

	_ = c.writeObject(writer, http.StatusOK, response)
}

// HandleDeleteCustomer handles deleting a customer.
func (c *Client) HandleDeleteCustomer(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := c.createContextWithTimeout(TimeoutHandleDeleteCustomer)
	defer cancel()

	type payload struct {
		Nik int `schema:"nik"`
	}
	p := &payload{}

	if err := json.NewDecoder(request.Body).Decode(p); err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to decode payload")
		return
	}

	err := c.DeleteUserCtx(ctx, p.Nik)
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to delete user")
		return
	}

	response := map[string]int{
		"nik": p.Nik,
	}

	_ = c.writeObject(writer, http.StatusOK, response)
}
