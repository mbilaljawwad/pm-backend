package main

import "github.com/mbilaljawwad/pm-backend/internal/server"

func main() {
	server := server.NewServer()
	server.Run()
}
