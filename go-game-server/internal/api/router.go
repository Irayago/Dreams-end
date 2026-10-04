package api

import (
	"fmt"

	"github.com/Irayago/Dreams-end/go-game-server/internal/hub"
	ws "github.com/coder/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Router struct {
	// define gin router here
	router *gin.Engine
	hubInt HubInterface
}

type HubInterface interface {
	ConnectClient(client *hub.Client) error
	DisconnectClient(client *hub.Client) error
	GetNumOfClients() int
	GetNumOfWorlds() int
}

func NewRouter(mainHub HubInterface) *Router {
	ginRouter := gin.Default()
	return &Router{
		router: ginRouter,
		hubInt: mainHub,
	}
}

// Router Operator methods below

func (r *Router) Run(port string) {
	r.RegisterRoutes()
	r.router.Run(port) //blocking call to start the HTTP server
}

func (r *Router) RegisterRoutes() {

	r.router.GET("/ws", r.websocketConnect)
	r.router.GET("/clients", r.getNumOfClients)
	r.router.GET("/worlds", r.getNumOfWorlds)

}

// Client facing APIs below

/*
When connecting a new WS client and after authorization:
1. GET request incoming to /ws endpoint
2. Router will call hub.ConnectClient() to add new client to Hub's client map if client isn't already connected
3. Hub will create a new World for the client if it doesn't exist, or connect the client to an existing World if it does exist
4. Hub will call world.AddPlayer() to add the player to the World
5. World will start sending messages to the client via the outboundBuffer channel
*/
func (r *Router) websocketConnect(c *gin.Context) {
	// handle WebSocket connection here

	// before accepting WS connection, do following:
	// 1. verify identity with JWT token; identity is defined by account name and playerId
	// 2. check for available worlds; if none, create a new world and assign to client
	wsConn, err := ws.Accept(c.Writer, c.Request, &ws.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Printf("Error accepting WebSocket connection: %v\n", err)
		c.Writer.Write(fmt.Appendf(nil, "Error accepting WebSocket connection: %v\n", err))
		return
	}

	// defer wsConn.CloseNow() // commenting out for now since we want to keep the connection open

	// create new Client
	client := hub.NewClient(wsConn, c.Request, generateClientId())
	err = r.hubInt.ConnectClient(client)
	if err != nil {
		fmt.Printf("Error connecting client to hub: %v\n", err)
		c.Writer.WriteHeader(500) // internal server error
		c.Writer.Write(fmt.Appendf(nil, "Error connecting client to hub: %v\n", err))
		return
	}

	// before http router exits goroutine, need to start client read and write pumps
	//go client.readPump(world)
	//go client.writePump()

}

func (r *Router) getNumOfClients(c *gin.Context) {
	if r.hubInt == nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write(fmt.Appendf(nil, "Error: Hub interface is nil\n"))
		return
	}
	c.JSON(200, gin.H{"numOfClients": r.hubInt.GetNumOfClients()})
}

func (r *Router) getNumOfWorlds(c *gin.Context) {
	if r.hubInt == nil {
		c.Writer.WriteHeader(500)
		c.Writer.Write(fmt.Appendf(nil, "Error: Hub interface is nil\n"))
		return
	}
	c.JSON(200, gin.H{"numOfWorlds": r.hubInt.GetNumOfWorlds()})
}

// Helper functions below

func generateClientId() string {
	id := uuid.NewString()
	return string(id)
}
