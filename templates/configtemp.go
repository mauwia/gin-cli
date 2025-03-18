package templates

// MainTemplate returns the basic Gin server code as a string.
// It replaces the placeholder with the provided module name.
func ConfigTemplate(moduleName string) string {
	return `package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds all config variables
type AppConfig struct {
   PORT string 
}

// LoadConfig loads environment variables from .env file into AppConfig struct
func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &AppConfig{
		PORT: getEnv("PORT", "8080"),
	}
}

// getEnv reads an environment variable and provides a fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

`
}
