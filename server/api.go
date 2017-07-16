package server

import (
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/heroku/go-with-me-app/database"
	"github.com/heroku/go-with-me-app/logger"
)

func StartAPIServer() {
	logger.Log.Infof("Running database migrations")
	err := database.RunDatabaseMigrations()
	if err != nil {
		logger.Log.Infof("database migration failed.")
	}

	logger.Log.Info("Starting `Service API")
	router := Router()
	handlerFunc := router.ServeHTTP

	server := negroni.New(negroni.NewRecovery())
	server.Use(httpStatLogger())
	server.UseHandlerFunc(handlerFunc)
	portInfo := ":" + os.Getenv("PORT")
	server.Run(portInfo)
}

func httpStatLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		next(rw, r)
		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime).Seconds() * 1000

		logger.Log.WithFields(logrus.Fields{
			"RequestTime":   startTime.Format(time.RFC3339),
			"ResponseTime":  responseTime.Format(time.RFC3339),
			"DeltaTime":     deltaTime,
			"RequestUrl":    r.URL.Path,
			"RequestMethod": r.Method,
			"RequestProxy":  r.RemoteAddr,
			"RequestSource": r.Header.Get("X-FORWARDED-FOR"),
		}).Debug("Http Logs")
	})
}
