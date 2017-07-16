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

func TestGetStatusReturnsErrorWhenBookingNotPresentInDatabase(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		bookingId := "111"
		expectedStatus := ""
		observedStatus, err := NewBookingRepository().GetBookingStatus(bookingId)
		assert.NotNil(t, err)
		assert.Equal(t, expectedStatus, observedStatus)
	})
}

func TestGetStatusReturnsErrorWhenBookingPresentInDatabase(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		expectedStatus := "created"
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"Status":      "CREATED",
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingRepository().CreateBooking(booking)
		assert.Nil(t, err)

		observedStatus, err := NewBookingRepository().GetBookingStatus(booking.BookingID)
		assert.Nil(t, err)
		assert.Equal(t, expectedStatus, observedStatus)
	})
}

func TestSetBookingStatusDriverNotFoundReturnsErrorWhenBookingNotPresentInDatabase(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		bookingId := "111"
		err := NewBookingRepository().SetBookingStatusDriverNotFound(bookingId)
		assert.NotNil(t, err)
	})
}

/* TODO;; add update status method
func TestSetBookingStatusDriverNotFoundReturnsErrorWhenBookingStatusIsNotCreated(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"Status":      "completed",
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingRepository().CreateBooking(booking)
		assert.Nil(t, err)

		err = NewBookingRepository().SetBookingStatusDriverNotFound(booking.BookingID)
		assert.NotNil(t, err)
	})
}
*/
func TestSetBookingStatusDriverNotFoundReturnsNilWhenBookingStatusIsCreated(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		err := NewBookingRepository().CreateBooking(booking)
		assert.Nil(t, err)

		err = NewBookingRepository().SetBookingStatusDriverNotFound(booking.BookingID)
		assert.Nil(t, err)

		expectedStatus := "driver_not_found"
		observedStatus, err := NewBookingRepository().GetBookingStatus(booking.BookingID)
		assert.Nil(t, err)
		assert.Equal(t, expectedStatus, observedStatus)
	})
}

func createBooking(data map[string]interface{}) *domain.Booking {
	return factories.BookingFactory.MustCreateWithOption(data).(*domain.Booking)
}
