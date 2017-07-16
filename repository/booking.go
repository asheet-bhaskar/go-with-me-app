package repository

import (
	"errors"
	"time"

	"github.com/heroku/go-with-me-app/appcontext"
	"github.com/heroku/go-with-me-app/domain"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type BookingRepository struct {
	db *sqlx.DB
}

const (
	createBookingQuery                  = `INSERT INTO bookings (booking_id, customer_id, driver_id, pick_up, destination, fare, status, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	getBookingStatusQuery               = `SELECT status FROM bookings WHERE booking_id=$1;`
	setBookingStatusDriverNotFoundQuery = `UPDATE bookings SET status='driver_not_found' WHERE booking_id=$1;`
)

func (br *BookingRepository) CreateBooking(booking *domain.Booking) error {
	now := time.Now()
	booking.BookingID = uuid.NewV4().String()
	booking.CreatedAt = now
	booking.UpdatedAt = now
	booking.Status = "created"
	_, err := br.db.Exec(createBookingQuery,
		booking.BookingID,
		booking.CustomerID,
		booking.DriverID,
		booking.PickUp,
		booking.Destination,
		booking.Fare,
		booking.Status,
		booking.CreatedAt,
		booking.UpdatedAt)
	return err
}

func (br *BookingRepository) GetBookingStatus(bookingId string) (string, error) {
	var status string
	err := br.db.Get(&status, getBookingStatusQuery, bookingId)
	if err != nil {
		return status, errors.New("Booking not found in the database")
	}
	return status, nil
}

func (br *BookingRepository) SetBookingStatusDriverNotFound(bookingId string) error {
	status, err := br.GetBookingStatus(bookingId)
	if err != nil {
		return err
	}

	if status != "created" {
		return errors.New("status can't be updated from " + status + "to driver not found")
	}

	_, err = br.db.Exec(setBookingStatusDriverNotFoundQuery, bookingId)
	if err != nil {
		return errors.New("failed to update the booking status to driver not found.")
	}
	return nil
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		db: appcontext.GetDB(),
	}
}
