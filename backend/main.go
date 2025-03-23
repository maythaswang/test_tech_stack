package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var server *APIServer = NewAPIServer(":8080")
	server.run();
}
