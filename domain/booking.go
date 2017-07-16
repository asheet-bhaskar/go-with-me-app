package domain

import (
	"errors"
	"time"
)

const (
	bookingValidationError = "required fields customer_id, pick_up, destination or fare are  missing "
)

type Booking struct {
	BookingID   string    `json:"booking_id"`
	CustomerID  string    `json:"customer_id"`
	DriverID    string    `json:"driver_id"`
	PickUp      string    `json:"pick_up"`
	Destination string    `json:"destination"`
	Fare        string    `json:"fare"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func Validate(booking Booking) error {
	if booking.CustomerID == "" || booking.PickUp == "" || booking.Destination == "" || booking.Fare == "" {
		return errors.New(bookingValidationError)
	}
	return nil
}
