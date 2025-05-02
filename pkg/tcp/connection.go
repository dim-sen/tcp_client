package tcp

import (
	"log"
	"net"
	"time"
)

const (
	maxRetries     = 3
	retryDelay     = 2 * time.Second
	ReconnectDelay = 5 * time.Second
)

type Connection struct {
	conn net.Conn
}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) Connect(addr string) error {
	var err error

	for retry := 0; retry < maxRetries; retry++ {
		c.conn, err = net.Dial("tcp", addr)
		if err == nil {
			// Set read and write deadlines
			c.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
			c.conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
			return nil
		}

		log.Printf("Connection attempt %d/%d failed: %v\n", retry+1, maxRetries, err)
		if retry < maxRetries-1 {
			time.Sleep(retryDelay)
		}
	}

	return err
}

func (c *Connection) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Connection) IsConnected() bool {
	return c.conn != nil
}

func (c *Connection) GetConn() net.Conn {
	return c.conn
}
