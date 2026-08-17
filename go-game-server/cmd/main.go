package main

import (
	"github.com/Irayago/Dreams-end/go-game-server/internal/api"
	"github.com/Irayago/Dreams-end/go-game-server/internal/hub"
)

const PORT = ":9999"

func main() {
	/*
		cfg := config.Load()
		db := store.Connect(cfg.Database)
	*/

	newHub := hub.NewHub()
	go newHub.Run()

	router := api.NewRouter(newHub) // pass cfg with server configs later
	router.Run(PORT)                //change to cfg.Port later

}
