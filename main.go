package main

import (
	_ "expvar"
	"os"

	"github.com/heroku/go-with-me-app/appcontext"
	"github.com/heroku/go-with-me-app/config"
	"github.com/heroku/go-with-me-app/console"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/urfave/cli"
)

func handleInitError() {
	if e := recover(); e != nil {
		logger.Log.Fatalf("Failed to load the app due to error : %s", e)
	}
}

func main() {
	defer handleInitError()

	config.Load()
	logger.SetupLogger()
	appcontext.Initiate()

	clientApp := cli.NewApp()
	clientApp.Name = "go-with-me-app"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Runs database migrations",
			Action: func(c *cli.Context) error {
				return console.RunDatabaseMigrations()
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
