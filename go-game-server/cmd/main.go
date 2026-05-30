package main

import (
	"fmt"

	hub "github.com/Irayago/Dreams-end/go-game-server/internal/hub"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error returned from main(): %v\n", err)
	}

}

func run() error {
	_, err := fmt.Println("Hello World!")
	if err != nil {
		fmt.Printf("Error returned from run(): %v\n", err)
	}

	newHub := hub.NewHub()
	newHub.Run()
	if err != nil {
		fmt.Printf("Error returned from newHub.Run(): %v\n", err)
	}

	return err
}
