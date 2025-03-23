package main

import (
	"log"
)

func main() {
	var server *APIServer = NewAPIServer(":8080")
	log.Printf("Starting up server")
	server.run()
}
