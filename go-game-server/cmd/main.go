package main

import (
	"fmt"
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
	return err
}

