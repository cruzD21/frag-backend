package main

import (
	"github.com/cruzD21/frag-backend/internal/server"
	"log"
)

func main() {

	app := server.CreateServer()

	if err := app.Run(); err != nil {
		log.Fatalf("app.Run ->  %v", err)
	}
}
