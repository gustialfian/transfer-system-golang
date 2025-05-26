package config

import (
	"os"
)

type Config struct {
	Port             string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

func LoadConfig() *Config {
	return &Config{
		Port:             getEnv("PORT", "8000"),
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDBName:   getEnv("POSTGRES_DBNAME", "transfer_system"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
