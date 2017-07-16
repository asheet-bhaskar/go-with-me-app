package service

import (
	"testing"

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

func TestCreateBookingReturnsErrorIfCustomerIdMissing(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
	})
}

func TestCreateBookingReturnsErrorIfPickUpMissing(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
	})
}

func TestCreateBookingReturnsErrorIfDestinationMissing(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"CustomerID": "111",
			"PickUp":     "100, 100",
			"Fare":       "10000.0"})
		err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
	})
}

func TestCreateBookingReturnsErrorIfFareMissing(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"PickUp":      "100, 100",
			"Destination": "111, 111"})
		err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
	})
}

func TestCreateBookingReturnsNilAndPersistBooking(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingService().CreateBooking(*booking)
		assert.Nil(t, err)
	})
}

func createBooking(data map[string]interface{}) *domain.Booking {
	return factories.BookingFactory.MustCreateWithOption(data).(*domain.Booking)
}
