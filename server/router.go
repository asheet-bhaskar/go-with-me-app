package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heroku/go-with-me-app/handler"
)

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)
	return router
}
