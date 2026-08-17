package hub

import (
	"net/http"

	"github.com/Irayago/Dreams-end/go-game-server/internal/world"
	ws "github.com/coder/websocket"
)

type Client struct {
	conn         *ws.Conn
	ipAddr       string
	connectionId string // UUID
	playerName   string
	worldId      *world.World
}

func NewClient(conn *ws.Conn, request *http.Request) *Client {
	return &Client{
		conn:         conn,
		ipAddr:       request.RemoteAddr,
		connectionId: "",
		playerName:   "",
		worldId:      nil,
	}
}

func (c *Client) SendData(data []byte) {

}

func (c *Client) GetPlayerName() string {
	return c.playerName
}
