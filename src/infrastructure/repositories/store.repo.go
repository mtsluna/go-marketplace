package repositories

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"google.golang.org/api/iterator"
	"log"
)

type StoreRepo struct {

}

const COLLECTION_STORE = "stores"

func (repo *StoreRepo) FindAll() [] contracts.Store {

	client := BaseRepo()
	var array [] contracts.Store
	iter := client.Collection(COLLECTION_STORE).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var store contracts.Store
		err = mapstructure.Decode(doc.Data(), &store)
		if err != nil {
			log.Fatalln(err)
		}
		store.Id = doc.Ref.ID
		array = append(array, store)
	}

	return array
}

func (repo *StoreRepo) FindById(id string) contracts.Store {

	client := BaseRepo()
	doc, err := client.Collection(COLLECTION_STORE).Doc(id).Get(ctx)
	var store contracts.Store
	err = mapstructure.Decode(doc.Data(), &store)
	if err != nil {
		log.Fatalln(err)
	}

	return store
}

func (repo *StoreRepo) Save(store contracts.Store) contracts.Store {

	client := BaseRepo()
	fmt.Println(store)
	doc, _, err := client.Collection(COLLECTION_STORE).Add(ctx, store)
	if err != nil {
		log.Fatalln(err)
	}
	store.Id = doc.ID

	return store
}

func (repo *StoreRepo) Update(id string, store contracts.Store) contracts.Store {

	client := BaseRepo()
	_, err := client.Collection(COLLECTION_STORE).Doc(id).Set(ctx, store)
	if err != nil {
		log.Fatalln(err)
	}
	store.Id = id

	return store
}