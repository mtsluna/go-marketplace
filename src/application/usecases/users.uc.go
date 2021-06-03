package usecases

import (
	"encoding/base64"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"github.com/mtsluna/go-market/src/infrastructure/repositories"
)

type UsersUc struct {
	repo repositories.UserRepo
}

func (uc *UsersUc) FindAll() [] contracts.User{

	var array = uc.repo.FindAll()

	return array

}

func (uc *UsersUc) FindById(id string) contracts.User {

	var user = uc.repo.FindById(id)

	return user

}

func (uc *UsersUc) Save(user contracts.User) contracts.User {

	//encript password
	codedPassword := base64.StdEncoding.EncodeToString([]byte(user.Password))

	//Assign encrypt password
	user.Password = codedPassword

	user = uc.repo.Save(user)

	return user

}

func (uc *UsersUc) Update(id string, user contracts.User) contracts.User {

	user = uc.repo.Update(id, user)

	return user

}

func (uc *UsersUc) Delete(id string) bool {

	deleted := uc.repo.Delete(id)

	return deleted

}