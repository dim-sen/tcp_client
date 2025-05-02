package message

import (
	"encoding/binary"
	"log"
	"net"
	"tcp_client/pkg/gps"
	"time"
)

const maxRetries = 3
const retryDelay = 2 * time.Second

func EncodeLocation(location gps.Location) []byte {
	// Convert location to bytes
	lat := uint64(location.Latitude * 1e6)  // Convert to microdegrees
	lon := uint64(location.Longitude * 1e6) // Convert to microdegrees
	timestamp := uint64(location.Timestamp.Unix())

	// Create message buffer
	message := make([]byte, 24) // 8 bytes each for lat, lon, and timestamp
	binary.BigEndian.PutUint64(message[0:8], lat)
	binary.BigEndian.PutUint64(message[8:16], lon)
	binary.BigEndian.PutUint64(message[16:24], timestamp)

	return message
}

func SendLocation(conn net.Conn, location gps.Location) error {
	message := EncodeLocation(location)

	// Create and send header
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(message)))

	// Retry sending the message
	for retry := 0; retry < maxRetries; retry++ {
		// Send header
		_, err := conn.Write(header)
		if err != nil {
			log.Printf("Error sending header (attempt %d/%d): %v\n", retry+1, maxRetries, err)
			if retry < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			return err
		}

		// Send message
		_, err = conn.Write(message)
		if err != nil {
			log.Printf("Error sending message (attempt %d/%d): %v\n", retry+1, maxRetries, err)
			if retry < maxRetries-1 {
				time.Sleep(retryDelay)
				continue
			}
			return err
		}

		// Reset write deadline
		conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
		return nil
	}

	return nil
}

func ReadResponse(conn net.Conn) ([]byte, error) {
	// Read response header
	responseHeader := make([]byte, 4)
	_, err := conn.Read(responseHeader)
	if err != nil {
		return nil, err
	}

	responseLength := binary.BigEndian.Uint32(responseHeader)
	response := make([]byte, responseLength)
	_, err = conn.Read(response)
	if err != nil {
		return nil, err
	}

	// Reset read deadline
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	return response, nil
}
