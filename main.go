package main

import (
	"fmt"

	"github.com/flashmob/go-guerrilla"
)

func main() {
	d := guerrilla.Daemon{}
	err := d.Start()

	if err == nil {
		fmt.Println("Server Started!")
	}

	for {
	}
}
