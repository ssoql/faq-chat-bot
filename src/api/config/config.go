package config

import "os"

const (
	LogLevel   = "info"
	appEnv     = "APP_ENV"
	production = "prod"
	port       = "8083"
)

func IsProduction() bool {
	return os.Getenv(appEnv) == production
}

func GetPort() string {
	return ":" + port
}
