package util

import (
	"fmt"
	"os"
)

// Exact mapping of expected ENV variables
type Config struct {
	APP_ENV     string
	API_PORT    string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
}

// Create a new ENV config
func NewConfig() *Config {
	return &Config{
		APP_ENV:     getEnv("APP_ENV", "development"),
		API_PORT:    getEnv("API_PORT", "8000"),
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "8765"),
		DB_USER:     getEnv("DB_USER_API", "apiuser"),
		DB_PASSWORD: getEnv("DB_USER_API_PASSWORD", "apiuser"),
		DB_DATABASE: getEnv("DB_DATABASE", "recipevault"),
	}
}

// Get database connection URL
func (c *Config) GetDBConnectionURL() string {
	return fmt.Sprintf(`
    host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASSWORD, c.DB_DATABASE)
}

// Get env value or fallback if not set
func getEnv(env string, fallback string) string {
	val, ok := os.LookupEnv(env)
	if !ok {
		return fallback
	}
	return val
}
