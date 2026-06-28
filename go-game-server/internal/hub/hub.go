package hub

import (
	"fmt"
	"net/http"

	"github.com/Irayago/Dreams-end/go-game-server/internal/world"
	l "github.com/Irayago/Dreams-end/go-game-server/pkg/logger"
	"github.com/google/uuid"

	ws "github.com/coder/websocket"
)

// Hub maintains the set of active clients. Hub registers new client connections to the Client map, and deregisters them when they disconnect.
/*
For-Select pattern will be used for handling register, unregister, and broadcast channels.
*/
type Hub struct {
	clients    map[string]*Client      // tracks active ws connections
	worlds     map[string]*world.World // tracks available worlds
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		worlds:     make(map[string]*world.World),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 256),
	}
}

func (h *Hub) Run() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := ws.Accept(w, r, nil)
		if err != nil {
			fmt.Printf("Error accepting WebSocket connection: %v\n", err)
			return
		}

		defer wsConn.CloseNow()

		// check if theres already a Client serving the same IP address. If so, then no need to create a new Clinet
		client := NewClient(wsConn, r)
		err = h.Register(client)
		if err != nil {
			fmt.Println(err)
		}

	})
}

func (h *Hub) Register(c *Client) error {
	if c == nil {
		return l.FormatError("Hub", "Client connection nil")
	}

	clientToRegister := <-h.register
	clientId := h.generateClientId()
	h.clients[clientId] = clientToRegister
	fmt.Printf("New client connected: %v\n", c.ipAddr)

	return nil
}

func (h Hub) generateClientId() string {
	id := uuid.NewString()
	return string(id)
}
