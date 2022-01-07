package config

import (
	"e-commerce/adapter/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer()Server{
	return Server{
		"5000",
		gin.Default(),
	}
}

func (s *Server) Run(){
	router := routes.ConfigRoutes(s.server)
	log.Println("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
