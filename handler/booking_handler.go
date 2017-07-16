package handler

import (
	"encoding/json"
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

		err = services.Booking.CreateBooking(booking)
		if err != nil {
			logger.Log.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
