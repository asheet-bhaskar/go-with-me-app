package service

import (
	"errors"
	"os"
	"strconv"

	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/repository"
)

type BookingService struct {
	repository *repository.BookingRepository
}

func (bs *BookingService) CreateBooking(booking domain.Booking) (domain.Booking, error) {
	err := domain.Validate(booking)
	if err != nil {
		return domain.Booking{}, err
	}
	err = bs.repository.CreateBooking(&booking)
	if err != nil {
		logger.Log.Info("failed to create the booking")
		return domain.Booking{}, err
	}
	return booking, nil
}

func (bs *BookingService) GetBookingStatus(bookingId string) (string, error) {
	if bookingId == "" {
		return "", errors.New("Must provide booking_id")
	}
	return bs.repository.GetBookingStatus(bookingId)
}

func (bs *BookingService) SetBookingStatusDriverNotFound(bookingId string) error {
	if bookingId == "" {
		return errors.New("Must provide booking_id")
	}
	return bs.repository.SetBookingStatusDriverNotFound(bookingId)
}

func (bs *BookingService) EstimateFare(distance string) (float64, error) {
	distanceFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return 0.0, errors.New("failed to parse the distance")
	}

	fareRateFloat, err := strconv.ParseFloat(os.Getenv("FARE_RATE"), 64)
	if err != nil {
		return 0.0, errors.New("failed to get the fare rate")
	}

	return distanceFloat * fareRateFloat, nil
}

func NewBookingService() *BookingService {
	return &BookingService{
		repository: repository.NewBookingRepository(),
	}
}
