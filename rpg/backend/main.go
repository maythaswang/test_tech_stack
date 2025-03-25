package main

func main() {
	var server *APIServer = NewAPIServer(":8080")
	server.run()
}
