package main

import (
	"encoding/binary"
	"log"
	"math/rand"
	"net"
	"time"
)

type GPSLocation struct {
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}

func generateRandomLocation() GPSLocation {
	// Generate random coordinates around a specific point (e.g., Jakarta)
	// Jakarta coordinates: -6.2088, 106.8456
	baseLat := -6.2088
	baseLon := 106.8456

	// Add small random variations
	lat := baseLat + (rand.Float64()-0.5)*0.01
	lon := baseLon + (rand.Float64()-0.5)*0.01

	return GPSLocation{
		Latitude:  lat,
		Longitude: lon,
		Timestamp: time.Now(),
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8181")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send GPS locations for 30 seconds
	stopTime := time.Now().Add(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for time.Now().Before(stopTime) {
		select {
		case <-ticker.C:
			location := generateRandomLocation()

			// Convert location to bytes
			lat := uint64(location.Latitude * 1e6)  // Convert to microdegrees
			lon := uint64(location.Longitude * 1e6) // Convert to microdegrees
			timestamp := uint64(location.Timestamp.Unix())

			// Create message buffer
			message := make([]byte, 24) // 8 bytes each for lat, lon, and timestamp
			binary.BigEndian.PutUint64(message[0:8], lat)
			binary.BigEndian.PutUint64(message[8:16], lon)
			binary.BigEndian.PutUint64(message[16:24], timestamp)

			// Create and send header
			header := make([]byte, 4)
			binary.BigEndian.PutUint32(header, uint32(len(message)))
			_, err = conn.Write(header)
			if err != nil {
				log.Println("Error sending header:", err)
				return
			}

			// Send message
			_, err = conn.Write(message)
			if err != nil {
				log.Println("Error sending message:", err)
				return
			}

			log.Printf("Sent location: Lat=%.6f, Lon=%.6f, Time=%s\n",
				location.Latitude,
				location.Longitude,
				location.Timestamp.Format(time.RFC3339))

			// Read response
			responseHeader := make([]byte, 4)
			_, err = conn.Read(responseHeader)
			if err != nil {
				log.Println("Error reading response header:", err)
				return
			}

			responseLength := binary.BigEndian.Uint32(responseHeader)
			response := make([]byte, responseLength)
			_, err = conn.Read(response)
			if err != nil {
				log.Println("Error reading response:", err)
				return
			}

			log.Printf("Server response: %s\n", response)
		}
	}
}
