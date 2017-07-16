package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	port     int
	logLevel string
	dbConfig *databaseConfig
}

var appConfig *Config

func Load() {
	viper.SetDefault("APP_PORT", "3000")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	appConfig = &Config{
		port:     getIntOrPanic("APP_PORT"),
		logLevel: fatalGetString("LOG_LEVEL"),
		dbConfig: newDatabaseConfig(),
	}

}

func Port() int {
	return appConfig.port
}

func LogLevel() string {
	return appConfig.logLevel
}

func AppConfig() *Config {
	return appConfig
}

func DatabaseConfig() *databaseConfig {
	return appConfig.dbConfig
}
