package repository

import (
	"testing"
	"time"

	"github.com/heroku/go-with-me-app/appcontext"
	"github.com/heroku/go-with-me-app/config"
	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/factories"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/stretchr/testify/assert"
)

func withCleanDB(block func()) {
	db := appcontext.GetDB()
	db.Exec("delete from bookings")
	block()
	db.Exec("delete from bookings")
	db.Close()
}

func TestCreateBookingWithUpdatedTimeFields(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{"BookingID": "B111",
			"CustomerID":  "111",
			"Status":      "CREATED",
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingRepository().CreateBooking(booking)

		assert.Nil(t, err)
		assert.NotEqual(t, time.Time{}, booking.CreatedAt)
		assert.NotEqual(t, time.Time{}, booking.UpdatedAt)

	})
}

func createBooking(data map[string]interface{}) *domain.Booking {
	return factories.BookingFactory.MustCreateWithOption(data).(*domain.Booking)
}
