package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/mtsluna/go-market/src/interface/handlers"
)

func StoreRouter(router *gin.RouterGroup) {

	router = router.Group("/stores")
	{
		storeHdlr := handlers.StoreHdlr{}

		router.GET("", storeHdlr.GetAll)
		router.GET("/:id", storeHdlr.GetById)
		router.POST("/", storeHdlr.Save)
		router.PUT("/:id", storeHdlr.Update)
	}

}