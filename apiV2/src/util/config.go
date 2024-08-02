package util

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string
}

// Create a new ENV config
func NewConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "8000"),
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

// Get formatted database connection string
func (c *Config) GetDBConnectionString() string {
	return fmt.Sprintf(`
    host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBDatabase)
}
