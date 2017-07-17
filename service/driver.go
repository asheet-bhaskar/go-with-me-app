package service

import (
	"errors"

	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/repository"
)

type DriverService struct {
	repository *repository.DriverRepository
}

func (ds *DriverService) CreateDriver(driver domain.Driver) (domain.Driver, error) {
	err := domain.ValidateDriver(driver)
	if err != nil {
		return domain.Driver{}, err
	}
	err = ds.repository.CreateDriver(&driver)
	if err != nil {
		logger.Log.Info("failed to create the driver")
		return domain.Driver{}, err
	}
	return driver, nil
}

func (ds *DriverService) UpdateDriverStatus(driverId string, status string) error {
	if driverId == "" || status == "" {
		return errors.New("Must provide driver_id and status")
	}
	return ds.repository.UpdateDriverStatus(driverId, status)
}

func (ds *DriverService) GetDriverStatus(driverId string) (string, error) {
	if driverId == "" {
		return "", errors.New("Must provide driver_id")
	}
	return ds.repository.GetDriverStatus(driverId)
}

func (ds *DriverService) CompleteTrip(driverId string, bookingId string) error {
	if driverId == "" || bookingId == "" {
		return errors.New("Must provide driver_id and bookingId")
	}
	return ds.repository.CompleteTrip(driverId, bookingId)
}

func (ds *DriverService) StartTrip(driverId string, bookingId string) error {
	if driverId == "" || bookingId == "" {
		return errors.New("Must provide driver_id and bookingId")
	}
	return ds.repository.StartTrip(driverId, bookingId)
}

func NewDriverService() *DriverService {
	return &DriverService{
		repository: repository.NewDriverRepository(),
	}
}
