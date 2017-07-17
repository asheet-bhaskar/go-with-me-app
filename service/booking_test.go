package service

import (
	"os"
	"strconv"
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

func TestSetBookingStatusDriverNotFoundReturnsErrorIfBookingIdIsEmpty(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	bookingId := ""
	err := NewBookingService().SetBookingStatusDriverNotFound(bookingId)
	assert.NotNil(t, err)
}

func TestSetBookingStatusDriverNotFoundReturnsNilIfBookingIdIsPresent(t *testing.T) {
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

		err = NewBookingService().SetBookingStatusDriverNotFound(res.BookingID)
		assert.Nil(t, err)
	})
}

func TestEstimateFareReturnsFareZeroAndErrorIfDistanceNotProvided(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	distance := ""
	estimatedFare := float64(0)
	observedFare, err := NewBookingService().EstimateFare(distance)
	assert.Equal(t, estimatedFare, observedFare)
	assert.NotNil(t, err)
}

func TestEstimateFareReturnsFareZeroAndErrorIfInvalidDistance(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	distance := "abc"
	estimatedFare := float64(0)
	observedFare, err := NewBookingService().EstimateFare(distance)
	assert.Equal(t, estimatedFare, observedFare)
	assert.NotNil(t, err)
}
func TestEstimateFareReturnsFareAndNilErrorIfValidDistance(t *testing.T) {
	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()
	distance := "9.5"
	os.Setenv("FARE_RATE", "2500")
	fareRateFloat, err := strconv.ParseFloat(os.Getenv("FARE_RATE"), 64)
	assert.Nil(t, err)

	estimatedFare := 9.5 * fareRateFloat
	observedFare, err := NewBookingService().EstimateFare(distance)
	assert.Equal(t, estimatedFare, observedFare)
	assert.Nil(t, err)
	os.Unsetenv("FARE_RATE")
}

func createBooking(data map[string]interface{}) *domain.Booking {
	return factories.BookingFactory.MustCreateWithOption(data).(*domain.Booking)
}
