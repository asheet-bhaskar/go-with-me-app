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
		res, err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Booking{}, res)
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
		res, err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Booking{}, res)
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
		res, err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Booking{}, res)
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
		res, err := NewBookingService().CreateBooking(*booking)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Booking{}, res)
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
		res, err := NewBookingService().CreateBooking(*booking)
		assert.Nil(t, err)
		assert.Equal(t, res.CustomerID, booking.CustomerID)
		assert.Equal(t, res.PickUp, booking.PickUp)
		assert.Equal(t, res.Destination, booking.Destination)
		assert.Equal(t, res.Fare, booking.Fare)
		assert.NotNil(t, res.CreatedAt)
		assert.NotNil(t, res.UpdatedAt)
	})
}

func TestGetBookingStatusReturnsErrorIfBookingIdIsEmpty(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	bookingId := ""
	expectedStatus := ""
	observedStatus, err := NewBookingService().GetBookingStatus(bookingId)
	assert.NotNil(t, err)
	assert.Equal(t, expectedStatus, observedStatus)
}

func TestGetBookingStatusReturnsNilIfBookingIdIsPresent(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	withCleanDB(func() {
		expectedStatus := "created"
		booking := createBooking(map[string]interface{}{
			"CustomerID":  "111",
			"PickUp":      "100, 100",
			"Destination": "111, 111",
			"Fare":        "10000.0"})
		res, err := NewBookingService().CreateBooking(*booking)
		assert.Nil(t, err)

		observedStatus, err := NewBookingService().GetBookingStatus(res.BookingID)
		assert.Nil(t, err)
		assert.Equal(t, expectedStatus, observedStatus)
	})
}

func createBooking(data map[string]interface{}) *domain.Booking {
	return factories.BookingFactory.MustCreateWithOption(data).(*domain.Booking)
}
