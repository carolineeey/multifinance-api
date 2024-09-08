package api

import (
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
	ctx, cancel := c.createContextWithTimeout(TimeoutHandleGetCustomer)
	defer cancel()

	type payload struct {
		Nik int `schema:"nik"`
	}
	p := &payload{}

	if err := c.decodeRequestSchema(writer, request, p); err != nil {
		return
	}

	customer, err := c.GetCustomerCtx(ctx, p.Nik)
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to get customer")
		return
	}

	_ = c.writeObject(writer, http.StatusOK, customer)
}

// HandleAddCustomer handles creating a new customer.
func (c *Client) HandleAddCustomer(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := c.createContextWithTimeout(TimeoutHandleCreateCustomer)
	defer cancel()

	payload := &SetUserRequest{}

	if err := c.decodeRequestJson(writer, request, payload); err != nil {
		return
	}

	err := c.CreateCustomerCtx(ctx, payload)
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

	if err := c.decodeRequestJson(writer, request, payload); err != nil {
		return
	}

	err := c.EditCustomerDataCtx(ctx, payload)
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

	if err := c.decodeRequestSchema(writer, request, p); err != nil {
		return
	}

	err := c.DeleteCustomerCtx(ctx, p.Nik)
	if err != nil {
		c.writeErrorObject(writer, err, http.StatusInternalServerError, "failed to delete user")
		return
	}

	response := map[string]int{
		"nik": p.Nik,
	}

	_ = c.writeObject(writer, http.StatusOK, response)
}
