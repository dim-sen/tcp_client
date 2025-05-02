package main

import (
	"log"
	"net"
	"tcp_client/pkg/gps"
	"tcp_client/pkg/message"
	"tcp_client/pkg/tcp"
	"time"
)

func main() {
	conn := tcp.NewConnection()
	defer conn.Close()

	// Send GPS locations for 30 seconds
	stopTime := time.Now().Add(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for time.Now().Before(stopTime) {
		select {
		case <-ticker.C:
			location := gps.GenerateRandomLocation()

			// Try to connect if not connected
			if !conn.IsConnected() {
				err := conn.Connect("localhost:8181")
				if err != nil {
					log.Println("Failed to connect to server:", err)
					time.Sleep(tcp.ReconnectDelay)
					continue
				}
				log.Println("Connected to server")
			}

			// Send location
			err := message.SendLocation(conn.GetConn(), location)
			if err != nil {
				log.Println("Failed to send location:", err)
				conn.Close()
				continue
			}

			log.Printf("Sent location: Lat=%.6f, Lon=%.6f, Time=%s\n",
				location.Latitude,
				location.Longitude,
				location.Timestamp.Format(time.RFC3339))

			// Read response
			response, err := message.ReadResponse(conn.GetConn())
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					log.Println("Read timeout, reconnecting...")
				} else {
					log.Println("Error reading response:", err)
				}
				conn.Close()
				continue
			}

			log.Printf("Server response: %s\n", response)
		}
	}
}
