package gps

import (
	"math/rand"
	"tcp_client/pkg/config"
	"time"
)

// Location represents a GPS position with timestamp
type Location struct {
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}

// GenerateRandomLocation creates a simulated GPS location
// This would be replaced by actual GPS device data in a real application
func GenerateRandomLocation() Location {
	// Add small random variations to base coordinates
	lat := config.BaseLatitude + (rand.Float64()-0.5)*config.Variation
	lon := config.BaseLongitude + (rand.Float64()-0.5)*config.Variation

	return Location{
		Latitude:  lat,
		Longitude: lon,
		Timestamp: time.Now(),
	}
}
