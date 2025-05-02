package gps

import (
	"math/rand"
	"time"
)

type Location struct {
	Latitude  float64
	Longitude float64
	Timestamp time.Time
}

func GenerateRandomLocation() Location {
	// Generate random coordinates around a specific point (e.g., Jakarta)
	// Jakarta coordinates: -6.2088, 106.8456
	baseLat := -6.2088
	baseLon := 106.8456

	// Add small random variations
	lat := baseLat + (rand.Float64()-0.5)*0.01
	lon := baseLon + (rand.Float64()-0.5)*0.01

	return Location{
		Latitude:  lat,
		Longitude: lon,
		Timestamp: time.Now(),
	}
}
