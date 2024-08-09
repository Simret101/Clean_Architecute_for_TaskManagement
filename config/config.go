package config

import (
	"os"
	"time"
)

// Configuration variables
var (
	SecretKey       string
	TokenExpiration time.Duration
)

func init() {
	// Load configuration from environment variables or use default values
	SecretKey = getEnv("SECRET_KEY", "mysecretkey")
	TokenExpiration = getEnvAsDuration("TOKEN_EXPIRATION", time.Minute*15)
}

// Helper functions

// getEnv retrieves environment variables or returns a default value if not found
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsDuration retrieves environment variables as a time.Duration
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
