package routes

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/mtsluna/go-market/src/interface/handlers"
)

func ImageRouter(router *gin.RouterGroup) {

	router = router.Group("/images")
	{
		imageHdlr := handlers.ImageHdlr{}

		router.POST("/upload", imageHdlr.Upload)
	}

}
