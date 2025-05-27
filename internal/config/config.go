// Package config provides functionality for loading application configuration
// from environment variables. It defines the Config struct, which holds
// configuration values such as server port and PostgreSQL connection details,
// and provides a function to load these values with sensible defaults.
package config

import (
	"os"
)

// Config struct which holds configuration values such as server port and PostgreSQL connection details,
type Config struct {
	Port             string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

// LoadConfig initializes and returns a pointer to a Config struct populated with
// values from environment variables.
func LoadConfig() *Config {
	return &Config{
		Port:             getEnv("APP_PORT", "8000"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresUser:     getEnv("POSTGRES_USER", ""),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", ""),
		PostgresDBName:   getEnv("POSTGRES_DBNAME", ""),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
