package network

import (
	"encoding/json"
	"net"
)

type Connection struct {
	conn   net.Conn
	encode *json.Encoder
	decode *json.Decoder
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn:   conn,
		encode: json.NewEncoder(conn),
		decode: json.NewDecoder(conn),
	}
}

func (c *Connection) Send(msg interface{}) error {
	return c.encode.Encode(msg)
}

func (c *Connection) Receive(msg interface{}) error {
	return c.decode.Decode(msg)
}

func (c *Connection) Close() error {
	return c.conn.Close()
}
