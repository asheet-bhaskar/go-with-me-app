package repository

import (
	"time"

	"github.com/heroku/go-with-me-app/appcontext"
	"github.com/heroku/go-with-me-app/domain"
	"github.com/jmoiron/sqlx"
)

type BookingRepository struct {
	db *sqlx.DB
}

const (
	createBookingQuery = `INSERT INTO bookings (booking_id, customer_id, driver_id, pick_up, destination, fare, status, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`
)

func (br *BookingRepository) CreateBooking(booking *domain.Booking) error {
	now := time.Now()
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

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		db: appcontext.GetDB(),
	}
}
