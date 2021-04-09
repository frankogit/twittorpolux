package main

import (
	"log"

	"github.com/frankogit/twittorpolux/bd"
	"github.com/frankogit/twittorpolux/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("No db connection")
		return
	}
	handlers.Manejadores()

}
