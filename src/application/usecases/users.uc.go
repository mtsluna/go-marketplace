package usecases

import (
	"encoding/base64"
	"github.com/mtsluna/go-market/src/domain/contracts"
	repo "github.com/mtsluna/go-market/src/infrastructure/repositories"
)

func UsersFindAll() [] contracts.User{

	var array = repo.FindAll()

	return array

}

func UsersFindById(id string) contracts.User {

	var user = repo.FindById(id)

	return user

}

func UsersSave(user contracts.User) contracts.User {

	//encript password
	codedPassword := base64.StdEncoding.EncodeToString([]byte(user.Password))

	//Assign encrypt password
	user.Password = codedPassword

	user = repo.Save(user)

	return user

}

func UsersUpdate(id string, user contracts.User) contracts.User {

	user = repo.Update(id, user)

	return user

}