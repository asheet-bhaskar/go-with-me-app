package domain

import (
	"errors"
	"time"
)

type Customer struct {
	CustomerID string
	Name       string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

const (
	customerValidationError = "required fields name or email are missing"
)

func ValidateCustomer(customer Customer) error {
	if customer.Name == "" || customer.Email == "" {
		return errors.New(customerValidationError)
	}
	return nil
}
