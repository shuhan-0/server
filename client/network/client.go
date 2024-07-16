package network

import (
	"encoding/json"
	"net"
)

type Client struct {
	conn    net.Conn
	encoder *json.Encoder
	decoder *json.Decoder
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:    conn,
		encoder: json.NewEncoder(conn),
		decoder: json.NewDecoder(conn),
	}, nil
}

func (c *Client) SendMessage(msgType string, payload interface{}) error {
	return c.encoder.Encode(map[string]interface{}{
		"type":    msgType,
		"payload": payload,
	})
}

func (c *Client) ReceiveMessage() (map[string]interface{}, error) {
	var msg map[string]interface{}
	err := c.decoder.Decode(&msg)
	return msg, err
}

func (c *Client) Close() error {
	return c.conn.Close()
}
