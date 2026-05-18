package hub

import (
	"fmt"
	"net/http"

	ws "github.com/coder/websocket"
)

// Hub maintains the set of active clients. Hub registers new client connections to the Client map, and deregisters them when they disconnect.
type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 256),
	}
}

func (h *Hub) Run() error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsConn, err := ws.Accept(w, r, nil)
		if err != nil {
			fmt.Printf("Error accepting WebSocket connection: %v\n", err)
			return
		}

		defer wsConn.CloseNow()

		client := NewClient(wsConn)
		h.clients[client] = true
		fmt.Printf("New client connected: %v\n", r.RemoteAddr)
	})
}
