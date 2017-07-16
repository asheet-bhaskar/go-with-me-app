package service

import (
	"errors"

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

func NewBookingService() *BookingService {
	return &BookingService{
		repository: repository.NewBookingRepository(),
	}
}
