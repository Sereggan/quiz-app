package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	DbAddress     string
}

func New() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		DbAddress:     getEnv("DB_ADDRESS", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
