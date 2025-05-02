package message

import (
	"encoding/binary"
	"tcp_client/pkg/gps"
)

// EncodeLocation converts a GPS location to a byte slice
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

// CreateMessageHeader creates a header for the message
func CreateMessageHeader(messageLength uint32) []byte {
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, messageLength)
	return header
}

// DecodeHeader parses a message header to get the message length
func DecodeHeader(header []byte) uint32 {
	return binary.BigEndian.Uint32(header)
}
