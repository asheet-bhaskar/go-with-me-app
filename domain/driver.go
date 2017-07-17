package domain

import (
	"errors"
	"time"
)

type Driver struct {
	DriverID  string
	Name      string
	Email     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	driverValidationError = "required fields name, email or status are missing"
)

func ValidateDriver(driver Driver) error {
	if driver.Name == "" || driver.Email == "" || driver.Status == "" {
		return errors.New(driverValidationError)
	}
	return nil
}
