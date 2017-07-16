package config

import "fmt"

type databaseConfig struct {
	host        string
	port        int
	username    string
	password    string
	name        string
	maxPoolSize int
}

func newDatabaseConfig() *databaseConfig {
	return &databaseConfig{
		host:        fatalGetString("host"),
		port:        getIntOrPanic("port"),
		name:        fatalGetString("database"),
		username:    getString("db_user"),
		password:    getString("password"),
		maxPoolSize: getIntOrPanic("pool"),
	}
}

func (dc *databaseConfig) DatabaseMaxPoolSize() int {
	return dc.maxPoolSize
}

func (dc *databaseConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=disable", dc.name, dc.username, dc.password, dc.host)
}

func (dc *databaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dc.username, dc.password, dc.host, dc.port, dc.name)
}
