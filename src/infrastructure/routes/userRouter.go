package routes

import (
	"github.com/gin-gonic/gin"
	usersHandler "github.com/mtsluna/go-market/src/interface/handlers"
)

func UserRouter(router *gin.RouterGroup) {

	router = router.Group("/users")
	{
		router.GET("", usersHandler.GetAll)
		router.GET("/:id", usersHandler.GetById)
		router.POST("/", usersHandler.Save)
		router.PUT("/:id", usersHandler.Update)
	}

}