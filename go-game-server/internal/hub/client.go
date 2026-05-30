package hub

import (
	"net/http"

	ws "github.com/coder/websocket"
)

type Client struct {
	conn *ws.Conn
	ipAddr string
}

func NewClient(conn *ws.Conn, request *http.Request) *Client {
	return &Client{
		conn: conn,
		ipAddr: request.RemoteAddr,
	}
}

func (*Client) sendData(data []byte) {
	
}

func (*Client) playerId() string {
	return ""
}
