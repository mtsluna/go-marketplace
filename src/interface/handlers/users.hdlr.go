package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mtsluna/go-market/src/application/usecases"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"log"
)

func GetAll(ctx *gin.Context) {

	users := usecases.UsersFindAll()
	var interfaceUsers = make([]interface{}, 0)
	for _, user := range users {
		interfaceUsers = append(interfaceUsers, user)
	}
	ctx.JSON(200, interfaceUsers)
}

func GetById(ctx *gin.Context) {

	id := ctx.Param("id")
	user := usecases.UsersFindById(id)
	ctx.JSON(200, user)
}

func Save(ctx *gin.Context) {

	var user contracts.User
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &user)
	user = usecases.UsersSave(user)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, user)
}

func Update(ctx *gin.Context) {

	id := ctx.Param("id")
	var user contracts.User
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &user)
	user = usecases.UsersUpdate(id, user)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, user)

}

