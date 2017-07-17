package repository

import (
	"time"

	"github.com/heroku/go-with-me-app/appcontext"
	"github.com/heroku/go-with-me-app/domain"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type CustomerRepository struct {
	db *sqlx.DB
}

const (
	createCustomerQuery = `INSERT INTO customers (customer_id, name, email, updated_at, created_at) VALUES($1, $2, $3, $4, $5);`
)

func (cr *CustomerRepository) CreateCustomer(customer *domain.Customer) error {
	now := time.Now()
	customer.CreatedAt = now
	customer.UpdatedAt = now
	customer.CustomerID = uuid.NewV4().String()
	_, err := cr.db.Exec(createCustomerQuery,
		customer.CustomerID,
		customer.Name,
		customer.Email,
		customer.CreatedAt,
		customer.UpdatedAt)
	return err
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: appcontext.GetDB(),
	}
}
