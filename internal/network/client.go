package network

import (
	"encoding/json"
	"net"
	"snake-game/pkg/protocol"
	"sync"
)

type Client struct {
	conn   net.Conn
	encode *json.Encoder
	decode *json.Decoder
	mutex  sync.Mutex
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		encode: json.NewEncoder(conn),
		decode: json.NewDecoder(conn),
	}, nil
}

func (c *Client) Send(msg protocol.Message) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.encode.Encode(msg)
}

func (c *Client) Receive() (protocol.Message, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var msg protocol.Message
	err := c.decode.Decode(&msg)
	return msg, err
}

func (c *Client) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.conn.Close()
}
