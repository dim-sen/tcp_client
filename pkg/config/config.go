package config

import (
	"os"
	"strconv"
	"time"
)

// Connection settings
var (
	// Server address
	ServerAddress string

	// Connection retry settings
	MaxRetries     int
	RetryDelay     time.Duration
	ReconnectDelay time.Duration

	// Timeout settings
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	// Data generation settings
	UpdateInterval time.Duration
)

// GPS base coordinates
var (
	BaseLatitude  float64
	BaseLongitude float64
	Variation     float64 // Maximum variation from base coordinates
)

// Initialize loads configuration from environment variables
func Initialize() {
	// Server settings
	ServerAddress = getEnv("TCP_SERVER_ADDRESS", "localhost:8181")

	// Connection retry settings
	MaxRetries = getEnvAsInt("TCP_MAX_RETRIES", 3)
	RetryDelay = getEnvAsDuration("TCP_RETRY_DELAY", 2*time.Second)
	ReconnectDelay = getEnvAsDuration("TCP_RECONNECT_DELAY", 5*time.Second)

	// Timeout settings
	ReadTimeout = getEnvAsDuration("TCP_READ_TIMEOUT", 60*time.Second)
	WriteTimeout = getEnvAsDuration("TCP_WRITE_TIMEOUT", 60*time.Second)

	// Data generation settings
	UpdateInterval = getEnvAsDuration("GPS_UPDATE_INTERVAL", 1*time.Second)

	// GPS coordinates
	BaseLatitude = getEnvAsFloat("GPS_BASE_LATITUDE", -6.2088)
	BaseLongitude = getEnvAsFloat("GPS_BASE_LONGITUDE", 106.8456)
	Variation = getEnvAsFloat("GPS_COORDINATES_VARIATION", 0.01)
}

// Helper function to get an environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to get an environment variable as an integer
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

// Helper function to get an environment variable as a float
func getEnvAsFloat(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}

// Helper function to get an environment variable as a duration
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}
