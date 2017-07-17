package handler

import (
	"encoding/json"
	"net/http"

	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/service"
)

func CreateCustomerHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer domain.Customer
		err := json.NewDecoder(r.Body).Decode(&customer)
		defer r.Body.Close()
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := services.Customer.CreateCustomer(customer)

		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		payload, err := json.Marshal(response)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
