package api

import (
	"context"
	"fmt"
	_ "github.com/carolineeey/multifinance-api/errnum"
	_ "github.com/go-sql-driver/mysql"
)

// GetCustomerCtx gets a customer.
func (c *Client) GetCustomerCtx(ctx context.Context, nik int) (customer *CustomerInfo, err error) {
	customer = &CustomerInfo{}
	err = c.Db.QueryRowContext(
		ctx,
		`SELECT full_name, legal_name, birth_place, birth_date, salary, ktp_photo_id, selfie_id FROM customers WHERE nik = ?`,
		nik,
	).Scan(
		&customer.FullName,
		&customer.LegalName,
		&customer.BirtPlace,
		&customer.BirthDate,
		&customer.Salary,
		&customer.KtpPhotoId,
		&customer.SelfieId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer info: %w", err)
	}

	return customer, nil
}

// CreateCustomerCtx creates a new customer.
func (c *Client) CreateCustomerCtx(ctx context.Context, customer *SetUserRequest) (err error) {
	_, err = c.Db.QueryContext(
		ctx,
		`INSERT INTO customers(nik, full_name, legal_name, birth_place, birth_date, salary) VALUES (?,?,?,?,FROM_UNIXTIME(?),?)`,
		customer.Nik,
		customer.FullName,
		customer.LegalName,
		customer.BirtPlace,
		customer.BirthDate,
		customer.Salary,
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

// EditCustomerDataCtx edits the data of a customer.
func (c *Client) EditCustomerDataCtx(ctx context.Context, customer *SetUserRequest) (err error) {
	_, err = c.Db.QueryContext(
		ctx,
		`UPDATE customers SET full_name = ?, legal_name = ?, birth_place = ?, birth_date = ?, salary = ? WHERE nik = ?`,
		customer.FullName,
		customer.LegalName,
		customer.BirtPlace,
		customer.BirthDate,
		customer.Salary,
		customer.Nik,
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

// DeleteCustomerCtx deletes a customer.
func (c *Client) DeleteCustomerCtx(ctx context.Context, nik int) (err error) {
	_, err = c.Db.QueryContext(
		ctx,
		`DELETE FROM customers WHERE nik = ?`,
		nik,
	)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
