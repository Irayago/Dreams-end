package api

import (
	"github.com/Irayago/Dreams-end/go-game-server/internal/hub"
	"github.com/gin-gonic/gin"
)

type Router struct {
	// define gin router here
	router *gin.Engine
}

type Hubinterface interface {
	ConnectClient(client *hub.Client)
	DisconnectClient(client *hub.Client)
}

func NewRouter() *Router {
	ginRouter := gin.Default()
	return &Router{
		router: ginRouter,
	}
}

/*
When connecting a new WS client and after authorization:
1. GET request incoming to /ws endpoint
2. Router will call hub.ConnectClient() to add new client to Hub's client map if client isn't already connected
3. Hub will create a new World for the client if it doesn't exist, or connect the client to an existing World if it does exist
4. Hub will call world.AddPlayer() to add the player to the World
5. World will start sending messages to the client via the outboundBuffer channel
*/

func (r *Router) Run(port string) {
	r.router.Run(port)
}

// pass Hub here
func (r *Router) RegisterRoutes(hub Hubinterface) {
}
