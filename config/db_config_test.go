package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConfig(t *testing.T) {
	configVars := map[string]string{
		"LOG_LEVEL": "debug",
		"APP_PORT":  "3000",
		"HOST":      "localhost",
		"PORT":      "5432",
		"DATABASE":  "db",
		"DB_USER":   "postgres",
		"PASSWORD":  "passs",
		"POOL":      "5",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	Load()
	assert.Equal(t, "postgres://postgres:passs@localhost:5432/db?sslmode=disable", DatabaseConfig().ConnectionURL())
}
