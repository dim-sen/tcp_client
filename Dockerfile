FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /gps_client cmd/tcp_app/main/main.go

# Create a minimal image
FROM alpine:latest

WORKDIR /

# Copy the binary from builder
COPY --from=builder /gps_client /gps_client

# Set default environment variables
ENV TCP_SERVER_ADDRESS=localhost:8181 \
    TCP_MAX_RETRIES=3 \
    TCP_RETRY_DELAY=2s \
    TCP_RECONNECT_DELAY=5s \
    TCP_READ_TIMEOUT=60s \
    TCP_WRITE_TIMEOUT=60s \
    GPS_UPDATE_INTERVAL=1s \
    GPS_BASE_LATITUDE=-6.2088 \
    GPS_BASE_LONGITUDE=106.8456 \
    GPS_COORDINATES_VARIATION=0.01

# Run the binary
CMD ["/gps_client"] 