package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mtsluna/go-market/src/application/usecases"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"log"
)

type UserHdlr struct {
	uc usecases.UsersUc
}

func (hdlr *UserHdlr) GetAll(ctx *gin.Context) {

	users := hdlr.uc.FindAll()
	var interfaceUsers = make([]interface{}, 0)
	for _, user := range users {
		interfaceUsers = append(interfaceUsers, user)
	}
	ctx.JSON(200, interfaceUsers)
}

func (hdlr *UserHdlr) GetById(ctx *gin.Context) {

	id := ctx.Param("id")
	user := hdlr.uc.FindById(id)
	ctx.JSON(200, user)
}

func (hdlr *UserHdlr) Save(ctx *gin.Context) {

	id := ctx.Param("id")
	var user contracts.User
	raw, err := ctx.GetRawData()
	err = json.Unmarshal(raw, &user)
	user = hdlr.uc.Save(id, user)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, user)
}

func (hdlr *UserHdlr) Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var user contracts.User
	raw, err := ctx.GetRawData()
	err = json.Unmarshal(raw, &user)
	user = hdlr.uc.Update(id, user)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, user)

}

func (hdlr *UserHdlr) Delete(ctx *gin.Context) {

	id := ctx.Param("id")
	deleted := hdlr.uc.Delete(id)

	if deleted == true {
		ctx.JSON(204, "")
	} else {
		ctx.JSON(404, "")
	}

}
