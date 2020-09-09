package config

import (
	"os"
)

type Config struct {
	Environment string
	Port        string
	Database    *Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func NewConfig() (*Config, error) {
	port := os.Getenv("PORT")

	// set default PORT if missing
	if port == "" {
		port = "8080"
	}

	return &Config{
		Environment: os.Getenv("ENV"),
		Port:        port,
		Database: &Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			DB:       os.Getenv("DATABASE_DB"),
			Password: os.Getenv("DATABASE_PASSWORD"),
		},
	}, nil
}

func (c *Config) GetPort() string {
	return ":" + c.Port
}
