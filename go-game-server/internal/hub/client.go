package hub

import (
	"net/http"

	"github.com/Irayago/Dreams-end/go-game-server/internal/world"
	ws "github.com/coder/websocket"
)

type Client struct {
	conn     *ws.Conn
	ipAddr   string
	playerId string
	worldId  *world.World
}

func NewClient(conn *ws.Conn, request *http.Request) *Client {
	return &Client{
		conn:   conn,
		ipAddr: request.RemoteAddr,
	}
}

func (c *Client) sendData(data []byte) {

}

func (c *Client) getplayerId() string {
	return c.playerId
}
