package message

import (
	"log"
	"tcp_client/pkg/config"
	"tcp_client/pkg/gps"
	"tcp_client/pkg/tcp"
	"time"
)

// SendLocation encodes and sends a GPS location to the server
func SendLocation(conn *tcp.Connection, location gps.Location) error {
	message := EncodeLocation(location)

	// Create header
	header := CreateMessageHeader(uint32(len(message)))

	// Retry sending the message
	for retry := 0; retry < config.MaxRetries; retry++ {
		// Reset deadlines before each attempt
		conn.ResetDeadlines()

		// Send header
		_, err := conn.GetConn().Write(header)
		if err != nil {
			log.Printf("Error sending header (attempt %d/%d): %v\n", retry+1, config.MaxRetries, err)
			if retry < config.MaxRetries-1 {
				time.Sleep(config.RetryDelay)
				continue
			}
			return err
		}

		// Send message
		_, err = conn.GetConn().Write(message)
		if err != nil {
			log.Printf("Error sending message (attempt %d/%d): %v\n", retry+1, config.MaxRetries, err)
			if retry < config.MaxRetries-1 {
				time.Sleep(config.RetryDelay)
				continue
			}
			return err
		}

		return nil
	}

	return nil
}

// ReadResponse reads and parses a response from the server
func ReadResponse(conn *tcp.Connection) ([]byte, error) {
	// Reset deadlines before reading
	conn.ResetDeadlines()

	// Read response header
	responseHeader := make([]byte, 4)
	_, err := conn.GetConn().Read(responseHeader)
	if err != nil {
		return nil, err
	}

	responseLength := DecodeHeader(responseHeader)
	response := make([]byte, responseLength)
	_, err = conn.GetConn().Read(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
