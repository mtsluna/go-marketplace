package repositories

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"google.golang.org/api/iterator"
	"log"
)

type UserRepo struct {

}

const COLLECTION = "users"

func (repo *UserRepo) FindAll() [] contracts.User {

	client := BaseRepo()
	var array [] contracts.User
	iter := client.Collection(COLLECTION).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var user contracts.User
		err = mapstructure.Decode(doc.Data(), &user)
		if err != nil {
			log.Fatalln(err)
		}
		user.Id = doc.Ref.ID
		user = removeDataFromUser(user)
		array = append(array, user)
	}

	return array
}

func (repo *UserRepo) FindById(id string) contracts.User {

	client := BaseRepo()
	doc, err := client.Collection(COLLECTION).Doc(id).Get(ctx)
	var user contracts.User
	err = mapstructure.Decode(doc.Data(), &user)
	if err != nil {
		log.Fatalln(err)
	}
	user = removeDataFromUser(user)

	return user
}

func (repo *UserRepo) Save(user contracts.User) contracts.User {

	client := BaseRepo()
	fmt.Println(user)
	doc, _, err := client.Collection(COLLECTION).Add(ctx, user)
	if err != nil {
		log.Fatalln(err)
	}
	user.Id = doc.ID
	user = removeDataFromUser(user)

	return user
}

func (repo *UserRepo) Update(id string, user contracts.User) contracts.User {

	client := BaseRepo()
	_, err := client.Collection(COLLECTION).Doc(id).Set(ctx, user)
	if err != nil {
		log.Fatalln(err)
	}
	user.Id = id
	user = removeDataFromUser(user)

	return user
}

func (repo *UserRepo) Delete(id string) bool{

	client := BaseRepo()
	_, err := client.Collection(COLLECTION).Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func removeDataFromUser(user contracts.User) contracts.User {
	user.Password = ""
	return user
}