package handler

import (
	"net/http"

	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/service"
)

func StartTripHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		driverId := r.URL.Query().Get("driver_id")
		bookingId := r.URL.Query().Get("booking_id")
		err := services.Driver.StartTrip(driverId, bookingId)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func CompleteTripHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		driverId := r.URL.Query().Get("driver_id")
		bookingId := r.URL.Query().Get("booking_id")
		err := services.Driver.CompleteTrip(driverId, bookingId)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
