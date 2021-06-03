package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	router "github.com/mtsluna/go-market/src/infrastructure/routes"
)

func main() {

	server := gin.Default()

	// Routers
	router.Router(server)

	server.Use(cors.Default())

	server.Run()
}


