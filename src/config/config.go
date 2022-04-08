package config

import "os"

const (
	LogLevel   = "info"
	appEnv     = "APP_ENV"
	production = "prod"
	port       = "8083"
	dbUsername = "APP_MYSQL_USERNAME"
	dbPassword = "APP_MYSQL_PASSWORD"
	dbHost     = "APP_MYSQL_HOST"
	dbPort     = "APP_MYSQL_PORT"
	dbSchema   = "APP_MYSQL_SCHEMA"
	esHosts    = "APP_ES_HOSTS"
)

func IsProduction() bool {
	return os.Getenv(appEnv) == production
}

func GetPort() string {
	return ":" + port
}

func GetDbUser() string {
	return os.Getenv(dbUsername)
}

func GetDbPassword() string {
	return os.Getenv(dbPassword)
}

func GetDbHost() string {
	return os.Getenv(dbHost)
}

func GetDbPort() string {
	return os.Getenv(dbPort)
}

func GetDbSchema() string {
	return os.Getenv(dbSchema)
}

func GetEsHosts() string {
	return os.Getenv(esHosts)
}
