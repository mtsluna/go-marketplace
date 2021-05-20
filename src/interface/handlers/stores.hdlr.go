package handlers

/*func StoresGetAll(ctx *gin.Context) {

	stores := usecases.UsersFindAll()
	var interfaceUsers = make([]interface{}, 0)
	for _, store := range stores {
		interfaceUsers = append(interfaceUsers, store)
	}
	ctx.JSON(200, interfaceUsers)
}

func StoresGetById(ctx *gin.Context) {

	id := ctx.Param("id")
	store := usecases.UsersFindById(id)
	ctx.JSON(200, store)
}

func StoresSave(ctx *gin.Context) {

	var store contracts.Store
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &store)
	store = usecases.StoresSave(store)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, store)
}

func StoresUpdate(ctx *gin.Context) {

	id := ctx.Param("id")
	var store contracts.Store
	raw, err :=  ctx.GetRawData()
	err = json.Unmarshal(raw, &store)
	store = usecases.UsersUpdate(id, store)

	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(201, store)

}*/


