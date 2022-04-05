package main

import (
	"backend/server"
	"log"

	"github.com/gin-contrib/cors"
)

func main() {
	srv, err := server.NewServer(
		server.OptServerPort("8080"),
		server.OptServerMiddleware(cors.Default()),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
