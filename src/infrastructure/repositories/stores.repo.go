package repositories

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"google.golang.org/api/iterator"
	"log"
)

const COLLECTION_STORES = "stores"

func StoreFindAll() [] contracts.User {

	client := BaseRepo()
	var array [] contracts.User
	iter := client.Collection(COLLECTION_STORES).Documents(CTX)
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

func StoreFindBy(id string) contracts.User {

	client := BaseRepo()
	doc, err := client.Collection(COLLECTION_STORES).Doc(id).Get(CTX)
	var user contracts.User
	err = mapstructure.Decode(doc.Data(), &user)
	if err != nil {
		log.Fatalln(err)
	}
	user = removeDataFromUser(user)

	return user
}

func StoreSave(user contracts.User) contracts.User {

	client := BaseRepo()
	fmt.Println(user)
	doc, _, err := client.Collection(COLLECTION_STORES).Add(CTX, user)
	if err != nil {
		log.Fatalln(err)
	}
	user.Id = doc.ID
	user = removeDataFromUser(user)

	return user
}

func try(){

}

func StoreUpdate(id string, user contracts.User) contracts.User {

	client := BaseRepo()
	_, err := client.Collection(COLLECTION_STORES).Doc(id).Set(CTX, user)
	if err != nil {
		log.Fatalln(err)
	}
	user.Id = id
	user = removeDataFromUser(user)

	return user
}
