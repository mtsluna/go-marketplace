package routes

import (
	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {

	dataRouter := server.Group("/api/marketplace")
	{
		UserRouter(dataRouter)
		StoreRouter(dataRouter)
		ProductRouter(dataRouter)
		ImageRouter(dataRouter)
	}
	
}