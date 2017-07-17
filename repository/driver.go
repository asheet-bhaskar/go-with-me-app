package repository

import (
	"errors"
	"time"

	"github.com/heroku/go-with-me-app/appcontext"
	"github.com/heroku/go-with-me-app/domain"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type DriverRepository struct {
	db *sqlx.DB
}

const (
	createDriverQuery       = `INSERT INTO drivers (driver_id, name, email, status, updated_at, created_at) VALUES($1, $2, $3, $4, $5, $6);`
	updateDriverStatusQurry = `UPDATE drivers set status=$1 WHERE driver_id=$2;`
	getDriverStatusQuery    = `SELECT status FROM bookings WHERE driver_id=$1;`
)

func (dr *DriverRepository) CreateDriver(driver *domain.Driver) error {
	now := time.Now()
	driver.CreatedAt = now
	driver.UpdatedAt = now
	driver.Status = "created"
	driver.DriverID = uuid.NewV4().String()
	_, err := dr.db.Exec(createDriverQuery,
		driver.DriverID,
		driver.Name,
		driver.Email,
		driver.Status,
		driver.CreatedAt,
		driver.UpdatedAt)
	return err
}

func (dr *DriverRepository) UpdateDriverStatus(driverId string, status string) error {
	_, err := dr.db.Exec(updateDriverStatusQurry, status, driverId)
	if err != nil {
		return errors.New("failed to update the driver status to " + status + "for driver " + driverId)
	}
	return nil
}

func (dr *DriverRepository) GetBookingStatus(driverId string) (string, error) {
	var status string
	err := dr.db.Get(&status, getDriverStatusQuery, driverId)
	if err != nil {
		return status, errors.New("driver not found in the database")
	}
	return status, nil
}

func NewDriverRepository() *DriverRepository {
	return &DriverRepository{
		db: appcontext.GetDB(),
	}
}
