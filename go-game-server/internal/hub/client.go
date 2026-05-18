package hub

import (
	ws "github.com/coder/websocket"
)

type Client struct {
	conn *ws.Conn
}

func NewClient(conn *ws.Conn) *Client {
	return &Client{
		conn: conn,
	}
}
