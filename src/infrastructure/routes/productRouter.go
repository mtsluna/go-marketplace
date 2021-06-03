package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/mtsluna/go-market/src/interface/handlers"
)

func ProductRouter(router *gin.RouterGroup) {

	router = router.Group("/products")
	{
		productHdlr := handlers.ProductsHdlr{}

		router.GET("", productHdlr.GetAll)
		router.GET("/by/store/:storeId", productHdlr.GetAllByStoreId)
		router.GET("/by/title/:title", productHdlr.GetAllByTitle)
		router.GET("/:id", productHdlr.GetById)
		router.POST("", productHdlr.Save)
		router.PUT("/:id", productHdlr.Update)
		router.DELETE("/:id", productHdlr.Delete)
	}

}
