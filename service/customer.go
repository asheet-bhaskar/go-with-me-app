package service

import (
	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/repository"
)

type CustomerService struct {
	repository *repository.CustomerRepository
}

func (cs *CustomerService) CreateCustomer(customer domain.Customer) (domain.Customer, error) {
	err := domain.ValidateCustomer(customer)
	if err != nil {
		return domain.Customer{}, err
	}
	err = cs.repository.CreateCustomer(&customer)
	if err != nil {
		logger.Log.Info("failed to create the customer")
		return domain.Customer{}, err
	}
	return customer, nil
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		repository: repository.NewCustomerRepository(),
	}
}
