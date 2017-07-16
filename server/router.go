package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heroku/go-with-me-app/handler"
	"github.com/heroku/go-with-me-app/service"
)

func Router(services *service.Services) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/v1/booking/create", handler.CreateBookingHandler(services)).Methods("POST")
	router.HandleFunc("/v1/booking/status", handler.GetBookingStatusHandler(services)).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
	return router
}
