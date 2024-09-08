package api

import (
	"context"
	"fmt"
	_ "github.com/carolineeey/multifinance-api/errnum"
	_ "github.com/go-sql-driver/mysql"
)

// CreateUserCtx creates a new user.
func (c *Client) CreateUserCtx(ctx context.Context, customer *SetUserRequest) (err error) {
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

// EditUserDataCtx edits the data of a user.
func (c *Client) EditUserDataCtx(ctx context.Context, customer *SetUserRequest) (err error) {
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

// DeleteUserCtx deletes a user.
func (c *Client) DeleteUserCtx(ctx context.Context, nik int) (err error) {
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
