package main

import (
	"fmt"
	"os"

	"github.com/flashmob/go-guerrilla"
)

func main() {
	d := guerrilla.Daemon{}
	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
		hName, _ := os.Hostname()
		fmt.Println("Host is", hName)
	}

	for {
	}
}
