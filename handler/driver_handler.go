package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/service"
)

func CreateDriverHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var driver domain.Driver
		err := json.NewDecoder(r.Body).Decode(&driver)
		defer r.Body.Close()
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := services.Driver.CreateDriver(driver)

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

func DriverStatusUpdateHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		driverId := r.URL.Query().Get("driver_id")
		status := r.URL.Query().Get("status")
		err := services.Driver.UpdateDriverStatus(driverId, status)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetDriverStatusHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		driverId := r.URL.Query().Get("driver_id")
		status, err := services.Driver.GetDriverStatus(driverId)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"status\": %v}", status)))
	}
}
