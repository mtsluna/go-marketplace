package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mtsluna/go-market/src/application/usecases"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"log"
)

type ProductsHdlr struct {
	uc usecases.ProductsUc
}

func (hdlr *ProductsHdlr) GetAll(ctx *gin.Context) {

	response := hdlr.uc.FindAll()
	var interfaceResponse = make([]interface{}, 0)
	for _, item := range response {
		interfaceResponse = append(interfaceResponse, item)
	}
	ctx.JSON(200, interfaceResponse)
}

func (hdlr *ProductsHdlr) GetAllByStoreId(ctx *gin.Context) {

	storeId := ctx.Param("storeId")

	response := hdlr.uc.FindAllByStoreId(storeId)
	var interfaceResponse = make([]interface{}, 0)
	for _, item := range response {
		interfaceResponse = append(interfaceResponse, item)
	}
	ctx.JSON(200, interfaceResponse)
}

func (hdlr *ProductsHdlr) GetAllByTitle(ctx *gin.Context) {

	title := ctx.Param("title")

	response := hdlr.uc.FindAllByTitle(title)
	var interfaceResponse = make([]interface{}, 0)
	for _, item := range response {
		interfaceResponse = append(interfaceResponse, item)
	}
	ctx.JSON(200, interfaceResponse)
}

func (hdlr *ProductsHdlr) GetById(ctx *gin.Context) {

	id := ctx.Param("id")
	item := hdlr.uc.FindById(id)
	ctx.JSON(200, item)
}

func (hdlr *ProductsHdlr) Save(ctx *gin.Context) {

	var item contracts.Product
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &item)
	item = hdlr.uc.Save(item)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, item)
}

func (hdlr *ProductsHdlr) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var item contracts.Product
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &item)
	item = hdlr.uc.Update(id, item)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, item)

}


