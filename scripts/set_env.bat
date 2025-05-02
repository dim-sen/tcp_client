@echo off
REM Configuration for TCP client application

REM Server settings
SET TCP_SERVER_ADDRESS=localhost:8181

REM Connection retry settings
SET TCP_MAX_RETRIES=3
SET TCP_RETRY_DELAY=2s
SET TCP_RECONNECT_DELAY=5s

REM Timeout settings
SET TCP_READ_TIMEOUT=60s
SET TCP_WRITE_TIMEOUT=60s

REM Data generation settings
SET GPS_UPDATE_INTERVAL=1s

REM GPS coordinates (Jakarta)
SET GPS_BASE_LATITUDE=-6.2088
SET GPS_BASE_LONGITUDE=106.8456
SET GPS_COORDINATES_VARIATION=0.01

echo Environment variables set successfully
echo Run "go run cmd/tcp_app/main/main.go" to start the application 