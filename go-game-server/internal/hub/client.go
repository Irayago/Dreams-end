package hub

import (
	"net/http"

	ws "github.com/coder/websocket"
)

type Client struct {
	conn         *ws.Conn
	ipAddr       string
	connectionId string // UUID
	playerName   string
	worldId      int
	inboundBuf   chan []MessageInterface
	outboundBuf  chan []MessageInterface
}

func NewClient(conn *ws.Conn, request *http.Request, connId string) *Client {
	return &Client{
		conn:         conn,
		ipAddr:       request.RemoteAddr,
		connectionId: connId,
		playerName:   "",
		worldId:      -1,
		inboundBuf:   make(chan []MessageInterface, 10),
		outboundBuf:  make(chan []MessageInterface, 10),
	}
}

type WorldInterface interface {
	GetWorldId() int
}

type MessageInterface interface {
	OpCode() int
	Payload() any
}

func (c *Client) SendData(data []byte) {

}

func (c *Client) GetPlayerName() string {
	return c.playerName
}

func (c *Client) SetPlayerName(name string) {
	c.playerName = name
}

func (c *Client) SavePlayerState() error {
	return nil
}

func (c *Client) GetWorldId() int {
	return c.worldId
}

func (c *Client) TearDownSession() error {
	return nil
}
