package domain

import (
	"errors"
	"time"
)

const (
	bookingValidationError = "required fields customer_id, pick_up, destination or fare are  missing "
)

type Booking struct {
	BookingID   string
	CustomerID  string
	DriverID    string
	PickUp      string
	Destination string
	Fare        string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Validate(booking Booking) error {
	if booking.CustomerID == "" || booking.PickUp == "" || booking.Destination == "" || booking.Fare == "" {
		return errors.New(bookingValidationError)
	}
	return nil
}
