package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/mtsluna/go-market/src/interface/handlers"
)

func UserRouter(router *gin.RouterGroup) {

	router = router.Group("/users")
	{
		userHdlr := handlers.UserHdlr{}

		router.GET("", userHdlr.GetAll)
		router.GET("/:id", userHdlr.GetById)
		router.POST("", userHdlr.Save)
		router.PUT("/:id", userHdlr.Update)
		router.DELETE("/:id", userHdlr.Delete)
	}

}