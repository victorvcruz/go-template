package config

import (
	"os"
)

type Database struct {
	Host     string
	User     string
	Password string
	Port     string
	DbName   string
	SSLMode  string
	TimeZone string
}

type Config struct {
	Database Database
}

func Load() (*Config, error) {
	return &Config{
		Database: Database{
			Host:     GetEnv("POSTGRES_HOST", "localhost"),
			User:     GetEnv("POSTGRES_USER", "postgres"),
			Password: GetEnv("POSTGRES_PASSWORD", "postgres"),
			Port:     GetEnv("POSTGRES_PORT", "5432"),
			DbName:   GetEnv("POSTGRES_DBNAME", "postgres"),
			SSLMode:  GetEnv("POSTGRES_SSLMODE", "disable"),
			TimeZone: GetEnv("POSTGRES_TIMEZONE", "UTC"),
		},
	}, nil
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
