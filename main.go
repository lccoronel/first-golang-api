package main

import (
	"main/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
