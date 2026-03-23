package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("port number failed to load or not available")
	}

	return &Config{
		Port: port,
	}, nil
}
