#!/bin/bash
# Configuration for TCP client application

# Server settings
export TCP_SERVER_ADDRESS=localhost:8181

# Connection retry settings
export TCP_MAX_RETRIES=3
export TCP_RETRY_DELAY=2s
export TCP_RECONNECT_DELAY=5s

# Timeout settings
export TCP_READ_TIMEOUT=60s
export TCP_WRITE_TIMEOUT=60s

# Data generation settings
export GPS_UPDATE_INTERVAL=1s

# GPS coordinates (Jakarta)
export GPS_BASE_LATITUDE=-6.2088
export GPS_BASE_LONGITUDE=106.8456
export GPS_COORDINATES_VARIATION=0.01

echo "Environment variables set successfully"
echo "Run 'go run cmd/tcp_app/main/main.go' to start the application" 