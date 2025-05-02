package main

import (
	"log"
	"math/rand"
	"tcp_client/pkg/client"
	"tcp_client/pkg/config"
	"time"
)

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize configuration from environment variables
	config.Initialize()

	// Log startup
	log.Printf("Starting GPS client, connecting to %s...", config.ServerAddress)

	// Create and start the GPS client
	client := client.NewGPSClient()
	client.Start()
}
