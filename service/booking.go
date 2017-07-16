package service

import (
	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/repository"
)

type bookingService struct {
	repository *repository.BookingRepository
}

func (bs *bookingService) CreateBooking(booking domain.Booking) error {
	err := domain.Validate(booking)
	if err != nil {
		return err
	}
	err = bs.repository.CreateBooking(&booking)
	if err != nil {
		logger.Log.Info("failed to create the booking")
		return err
	}
	return nil
}

func NewBookingService() *bookingService {
	return &bookingService{
		repository: repository.NewBookingRepository(),
	}
}
