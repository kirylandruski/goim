package network

import (
	"net"
	"unsafe"
	"encoding/binary"
	"context"
)

type Connection struct {
	conn     net.Conn
	doneChan chan bool
}

func NewConnection(ctx context.Context, netConn net.Conn) *Connection {
	c := &Connection{}
	c.conn = netConn

	c.doneChan = make(chan bool)

	go func() {
		for {
			select {
			case <-ctx.Done():
				c.Close()
				break
			case <-c.doneChan:
				return
			}
		}
	}()

	return c
}

func (c *Connection) Read() ([]byte, error) {
	metaBuf := make([]byte, unsafe.Sizeof(uint32(0)))
	if _, err := c.conn.Read(metaBuf); err != nil {
		c.Close()
		return nil, err
	}

	l := binary.BigEndian.Uint32(metaBuf)
	buf := make([]byte, l)

	if _, err := c.conn.Read(buf); err != nil {
		c.Close()
		return nil, err
	}

	return buf, nil
}

func (c *Connection) Write(payload []byte) error {
	lenBuf := make([]byte, unsafe.Sizeof(uint32(0)))
	binary.BigEndian.PutUint32(lenBuf, uint32(len(payload)))
	if _, err := c.conn.Write(lenBuf); err != nil {
		c.Close()
		return err
	}

	if _, err := c.conn.Write(payload); err != nil {
		c.Close()
		return err
	}

	return nil
}

func (c *Connection) Close() {
	c.doneChan <- true
	c.conn.Close()
}
