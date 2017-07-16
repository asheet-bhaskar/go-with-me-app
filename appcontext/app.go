package appcontext

import (
	"os"

	"github.com/heroku/go-with-me-app/config"
	"github.com/heroku/go-with-me-app/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type appContext struct {
	db *sqlx.DB
}

var context *appContext

type appContextError struct {
	Error error
}

func panicIfError(err error, werr error) {
	if err != nil {
		panic(appContextError{werr})
	}
}

func initDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Log.Fatalf("failed to load the database: %s", err)
	}

	if err = db.Ping(); err != nil {
		logger.Log.Fatalf("ping to the database host failed: %s", err)
	}

	db.SetMaxOpenConns(config.DatabaseConfig().DatabaseMaxPoolSize())
	return db
}

func Initiate() {
	db := initDB()
	context = &appContext{
		db: db,
	}
}

func GetDB() *sqlx.DB {
	return context.db
}
