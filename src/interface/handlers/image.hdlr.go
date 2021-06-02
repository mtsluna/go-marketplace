package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mtsluna/go-market/src/application/usecases"
)

type ImageHdlr struct {
	uc usecases.ImageUc
}

func (hdlr *ImageHdlr) Upload(ctx *gin.Context) {

	form, _ := ctx.MultipartForm()
	file := form.File["image"][0]

	image := hdlr.uc.Upload(file)

	ctx.JSON(200, image)
}