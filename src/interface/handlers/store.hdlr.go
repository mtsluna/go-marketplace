package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mtsluna/go-market/src/application/usecases"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"log"
)

type StoreHdlr struct {
	uc usecases.StoreUc
}

func (hdlr *StoreHdlr) GetAll(ctx *gin.Context) {

	response := hdlr.uc.FindAll()
	var interfaceResponse = make([]interface{}, 0)
	for _, item := range response {
		interfaceResponse = append(interfaceResponse, item)
	}
	ctx.JSON(200, interfaceResponse)
}

func (hdlr *StoreHdlr) GetById(ctx *gin.Context) {

	id := ctx.Param("id")
	item := hdlr.uc.FindById(id)
	ctx.JSON(200, item)
}

func (hdlr *StoreHdlr) Save(ctx *gin.Context) {

	var item contracts.Store
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &item)
	item = hdlr.uc.Save(item)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, item)
}

func (hdlr *StoreHdlr) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var item contracts.Store
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &item)
	item = hdlr.uc.Update(id, item)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, item)

}

func (hdlr *StoreHdlr) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	deleted := hdlr.uc.Delete(id)

	if deleted == true {
		ctx.JSON(204, "")
	} else {
		ctx.JSON(404, "")
	}

}

