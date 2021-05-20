package main

import (
	"github.com/gin-gonic/gin"
	router "github.com/mtsluna/go-market/src/infrastructure/routes"
)

func main() {

	server := gin.Default()

	// Routers
	router.Router(server)

	server.Run()
}


