package tcp

import (
	"log"
	"net"
	"tcp_client/pkg/config"
	"time"
)

// Connection manages a TCP connection with auto-reconnect capabilities
type Connection struct {
	conn net.Conn
}

// NewConnection creates a new TCP connection manager
func NewConnection() *Connection {
	return &Connection{}
}

// Connect establishes a TCP connection to the specified address with retry logic
func (c *Connection) Connect(addr string) error {
	var err error

	for retry := 0; retry < config.MaxRetries; retry++ {
		c.conn, err = net.Dial("tcp", addr)
		if err == nil {
			// Set read and write deadlines
			c.conn.SetReadDeadline(time.Now().Add(config.ReadTimeout))
			c.conn.SetWriteDeadline(time.Now().Add(config.WriteTimeout))
			return nil
		}

		log.Printf("Connection attempt %d/%d failed: %v\n", retry+1, config.MaxRetries, err)
		if retry < config.MaxRetries-1 {
			time.Sleep(config.RetryDelay)
		}
	}

	return err
}

// Close closes the current connection if active
func (c *Connection) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// IsConnected checks if connection is currently established
func (c *Connection) IsConnected() bool {
	return c.conn != nil
}

// GetConn returns the underlying net.Conn object
func (c *Connection) GetConn() net.Conn {
	return c.conn
}

// ResetDeadlines updates the read/write deadlines on the connection
func (c *Connection) ResetDeadlines() {
	if c.conn != nil {
		c.conn.SetReadDeadline(time.Now().Add(config.ReadTimeout))
		c.conn.SetWriteDeadline(time.Now().Add(config.WriteTimeout))
	}
}
