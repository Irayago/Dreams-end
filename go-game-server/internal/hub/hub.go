package hub

import (
	"fmt"
	"net/http"

	"github.com/Irayago/Dreams-end/go-game-server/internal/world"
	"github.com/google/uuid"

	ws "github.com/coder/websocket"
)

// Hub maintains the set of active clients. Hub connects new client connections to the Client map, and disconnects from the map when they disconnect.
// Hub needs to enforce one WS connection per client
/*
For-Select pattern will be used for handling connect, disconnect, and broadcast channels.
*/
type Hub struct {
	clients    map[string]*Client      // tracks active ws connections; key is clientId, value is Client struct ptr
	worlds     map[string]*world.World // tracks available worlds; key is worldId, value is World struct ptr
	connect    chan *Client            // renamed from register to connect for only tracking active WS connections
	disconnect chan *Client            // same as connect but for disconnecting clients
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		worlds:     make(map[string]*world.World),
		connect:    make(chan *Client, 100),
		disconnect: make(chan *Client, 100),
		broadcast:  make(chan []byte, 200),
	}
}

func (h *Hub) Run() {
	/*
		// define Hub API endpoints
		http.HandleFunc("/ws", h.webSocketHandler)

		// define server config; need to move to config.go later
		httpServer := &http.Server{
			Addr:         ":9999",
			Handler:      nil,
			ReadTimeout:  0,
			WriteTimeout: 0,
			IdleTimeout:  0,
		}

		err := httpServer.ListenAndServe()
		if err != nil {
			fmt.Printf("Error thrown from httpServer.ListenAndServe(): %v\n", err)
		}
	*/
	for {
		select {
		case client := <-h.connect:
			h.Connect(client)
		case client := <-h.disconnect:
			h.Disconnect(client)
		case message := <-h.broadcast:
			for _, client := range h.clients {
				client.SendData(message)
			}
		}
	}
}

// handler for managing weboscket connection to a client; gets passed to api.NewRouter()
func (h *Hub) webSocketHandler(w http.ResponseWriter, r *http.Request) {

	// before accepting WS connection, do following:
	// 1. verify identity with JWT token; identity is defined by account name and playerId
	// 2. check for available worlds; if none, create a new world and assign to client

	wsConn, err := ws.Accept(w, r, nil)
	if err != nil {
		fmt.Printf("Error accepting WebSocket connection: %v\n", err)
		w.WriteHeader(500) // internal server error
		w.Write(fmt.Appendf(nil, "Error accepting WebSocket connection: %v\n", err))
		return
	}

	defer wsConn.CloseNow()

	client := NewClient(wsConn, r)
	client.connectionId = h.generateClientId() // generate unique client connection ID
	h.connect <- client                        // send new client to connect channel for Hub to track

	// before http router exits goroutine, need to start client read and write pumps
	//go client.readPump(world)
	//go client.writePump()
}

func (h *Hub) Disconnect(c *Client) error {
	if c == nil {
		return fmt.Errorf("Hub: Client connection nil")
	}

	if _, ok := h.clients[c.connectionId]; !ok {
		return fmt.Errorf("Hub: Client connection not found")
	}

	delete(h.clients, c.connectionId)
	fmt.Printf("Client disconnected:\nIP: %v\nConnection ID: %v\n", c.ipAddr, c.connectionId)

	return nil
}

func (h *Hub) Connect(c *Client) error {
	if c == nil {
		return fmt.Errorf("Hub: Client connection nil")
	}

	// could happen if theres 2 active Clients with the same connectionId; very rare if due to UUID collision
	if _, ok := h.clients[c.connectionId]; ok {
		return fmt.Errorf("Hub: Client %v connection already exists", c.connectionId)
	}

	h.clients[c.connectionId] = c
	fmt.Printf("New client connected:\nIP: %v\nConnection ID: %v\n", c.ipAddr, c.connectionId)

	return nil
}

func (h Hub) generateClientId() string {
	id := uuid.NewString()
	return string(id)
}

// client interfaces

type ClientInterface interface {
}
