package util

import (
	"os"
)

type Config struct {
	APIPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string
}

// Create a new ENV config
func NewConfig() *Config {
	return &Config{
		APIPort:    getEnv("API_PORT", "8000"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "8765"),
		DBUser:     getEnv("DB_USER_API", "apiuser"),
		DBPassword: getEnv("DB_USER_API_PASSWORD", "apiuser"),
		DBDatabase: getEnv("DB_DATABASE", "recipehub"),
	}
}

// Get env value or fallback if not set
func getEnv(env string, fallback string) string {
	val, ok := os.LookupEnv(env)
	if !ok {
		return fallback
	}
	return val
}
