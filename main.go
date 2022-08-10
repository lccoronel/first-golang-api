package main

import (
	"main/database"
	"main/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
