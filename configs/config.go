package configs

import (
	"fmt"
	"os"
	"strconv"
)

// Config stores all configuration of the application
type Config struct {
	// Database configurations
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSchema   string

	// Keycloak configurations
	KeycloakBaseURL  string
	KeycloakRealm    string
	KeycloakClientID string

	// Server configurations
	Port int
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %v", err)
	}

	return &Config{
		// Database configurations
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "keycloak"),
		DBSchema:   getEnv("DB_SCHEMA", "public"),

		// Keycloak configurations
		KeycloakBaseURL:  getEnv("KEYCLOAK_BASE_URL", "http://localhost:8080"),
		KeycloakRealm:    getEnv("KEYCLOAK_REALM", "fitnis"),
		KeycloakClientID: getEnv("KEYCLOAK_CLIENT_ID", "patient-service"),

		// Server configurations
		Port: port,
	}, nil
}

// Helper function to get environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
