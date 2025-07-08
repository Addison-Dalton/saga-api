package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load initializes the configuration by loading the .env file.
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, reading from environment.")
	}
}

// Get retrieves a configuration value from the environment, fatally exiting if not set.
func Get(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
