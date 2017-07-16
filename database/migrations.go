package console

import (
	"fmt"
	"strings"

	"github.com/heroku/go-with-me-app/config"

	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
)

const dbMigrationsPath = "./migrations"

func RunDatabaseMigrations() error {
	allErrors, ok := migrate.UpSync(config.DatabaseConfig().ConnectionURL(), dbMigrationsPath)
	if !ok {
		return joinErrors(allErrors)
	}

	fmt.Println("Migration successful")

	return nil
}

func joinErrors(errors []error) error {
	var errorMsgs []string
	for _, err := range errors {
		errorMsgs = append(errorMsgs, err.Error())
	}

	return fmt.Errorf(strings.Join(errorMsgs, ","))
}
