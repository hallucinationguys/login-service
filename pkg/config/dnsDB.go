package config

import "fmt"

func DsnDB(config *Config) string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	return dsn
}
