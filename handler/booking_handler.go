package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heroku/go-with-me-app/domain"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/heroku/go-with-me-app/service"
)

func CreateBookingHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var booking domain.Booking
		err := json.NewDecoder(r.Body).Decode(&booking)
		defer r.Body.Close()
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := services.Booking.CreateBooking(booking)

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

func GetBookingStatusHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookingId := r.URL.Query().Get("booking_id")
		status, err := services.Booking.GetBookingStatus(bookingId)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"status\": %v}", status)))
	}
}

func DriverNotFoundHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookingId := r.URL.Query().Get("booking_id")
		err := services.Booking.SetBookingStatusDriverNotFound(bookingId)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func handlerEstimateFareHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		distance := r.URL.Query().Get("distance")
		fare, err := services.Booking.EstimateFare(distance)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf("{\"estimate_fare\": %v}", fare)))
	}

}
