package server

import (
	"log"
	"main/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3333",
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	router := routes.ConfigRoutes(server.server)

	log.Print("Server is running on port ", server.port)
	log.Fatal(router.Run(":" + server.port))
}
