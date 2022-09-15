package database

import "fmt"

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.UserName, config.Password, config.Host, config.Port, config.Database)
	return connectionString
}
