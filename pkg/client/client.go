package client

import (
	"log"
	"net"
	"tcp_client/pkg/config"
	"tcp_client/pkg/gps"
	"tcp_client/pkg/message"
	"tcp_client/pkg/tcp"
	"time"
)

// GPSClient manages the GPS data sending client
type GPSClient struct {
	conn *tcp.Connection
}

// NewGPSClient creates a new GPS data client
func NewGPSClient() *GPSClient {
	return &GPSClient{
		conn: tcp.NewConnection(),
	}
}

// Start begins the client operation, sending GPS data continuously
func (c *GPSClient) Start() {
	defer c.conn.Close()

	// Create a ticker for periodic GPS updates
	ticker := time.NewTicker(config.UpdateInterval)
	defer ticker.Stop()

	// Run indefinitely
	for {
		select {
		case <-ticker.C:
			c.processGPSData()
		}
	}
}

// processGPSData handles a single GPS data point
func (c *GPSClient) processGPSData() {
	// In a real application, this would get data from an actual GPS device
	// For now, we're using simulated data
	location := gps.GenerateRandomLocation()

	// Try to connect if not connected
	if !c.conn.IsConnected() {
		if err := c.connect(); err != nil {
			return
		}
	}

	// Send location
	if err := c.sendLocation(location); err != nil {
		return
	}

	// Read response
	c.readResponse()
}

// connect establishes connection to the server
func (c *GPSClient) connect() error {
	err := c.conn.Connect(config.ServerAddress)
	if err != nil {
		log.Println("Failed to connect to server:", err)
		time.Sleep(config.ReconnectDelay)
		return err
	}
	log.Println("Connected to server")
	return nil
}

// sendLocation sends GPS location to the server
func (c *GPSClient) sendLocation(location gps.Location) error {
	err := message.SendLocation(c.conn, location)
	if err != nil {
		log.Println("Failed to send location:", err)
		c.conn.Close()
		return err
	}

	log.Printf("Sent location: Lat=%.6f, Lon=%.6f, Time=%s\n",
		location.Latitude,
		location.Longitude,
		location.Timestamp.Format(time.RFC3339))
	return nil
}

// readResponse reads and processes server response
func (c *GPSClient) readResponse() {
	response, err := message.ReadResponse(c.conn)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Println("Read timeout, reconnecting...")
		} else {
			log.Println("Error reading response:", err)
		}
		c.conn.Close()
		return
	}

	log.Printf("Server response: %s\n", response)
}
