package hub

import (
	"fmt"

	"sync"

	"github.com/Irayago/Dreams-end/go-game-server/internal/world"
)

// Hub maintains the set of active clients. Hub connects new client connections to the Client map, and disconnects from the map when they disconnect.
// Hub needs to enforce one WS connection per client
/*
For-Select pattern will be used for handling connect, disconnect, and broadcast channels.
*/
type Hub struct {
	mu         sync.Mutex
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
			h.RegisterClient(client)
			if len(h.worlds) == 0 {
				// if no worlds exist, wait for a new world or put client in queue
			}
		case client := <-h.disconnect:
			h.DeregisterClient(client)
		case message := <-h.broadcast:
			for _, client := range h.clients {
				client.SendData(message)
			}
		}
	}
}

func (h *Hub) DeregisterClient(c *Client) error {
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

func (h *Hub) RegisterClient(c *Client) error {
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

// need to expose following Hub methods as client facing APIs in router.go:

func (h *Hub) ConnectClient(c *Client) error {
	h.connect <- c
	return nil
}

func (h *Hub) DisconnectClient(c *Client) error {
	h.disconnect <- c
	return nil
}

func (h *Hub) GetNumOfClients() int16 {
	h.mu.Lock()
	defer h.mu.Unlock()
	return int16(len(h.clients))
}

func (h *Hub) GetNumOfWorlds() int16 {
	h.mu.Lock()
	defer h.mu.Unlock()
	return int16(len(h.worlds))
}

// client interfaces

type ClientInterface interface {
}
