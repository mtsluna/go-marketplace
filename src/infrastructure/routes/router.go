package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mtsluna/go-market/src/interface/middlewares"
)

func Router(server *gin.Engine) {

	dataRouter := server.Group("/api/marketplace")
	{

		dataRouter.Use(middlewares.AllowCors)

		UserRouter(dataRouter)
		StoreRouter(dataRouter)
		ProductRouter(dataRouter)
		ImageRouter(dataRouter)
	}
	
}