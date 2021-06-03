package repositories

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/mtsluna/go-market/src/domain/contracts"
	"google.golang.org/api/iterator"
	"log"
)

type ProductsRepo struct {

}

const COLLECTION_PRODUCTS = "products"

func (repo *ProductsRepo) FindAll() [] contracts.Product {

	client := BaseRepo()
	var array [] contracts.Product
	iter := client.Collection(COLLECTION_PRODUCTS).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var item contracts.Product
		err = mapstructure.Decode(doc.Data(), &item)
		if err != nil {
			log.Fatalln(err)
		}
		item.Id = doc.Ref.ID
		array = append(array, item)
	}

	return array
}

func (repo *ProductsRepo) FindAllByStoreId(id string) [] contracts.Product {

	client := BaseRepo()
	var array [] contracts.Product
	iter := client.CollectionGroup(COLLECTION_PRODUCTS).Where("storeId", "==", ""+id+"").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var item contracts.Product
		err = mapstructure.Decode(doc.Data(), &item)
		if err != nil {
			log.Fatalln(err)
		}
		item.Id = doc.Ref.ID
		array = append(array, item)
	}

	return array
}

func (repo *ProductsRepo) FindAllByTitle(title string) [] contracts.Product {

	client := BaseRepo()
	var array [] contracts.Product
	iter := client.CollectionGroup(COLLECTION_PRODUCTS).Where("title", ">=", title).Where("title", "<=", title+"\uf8ff").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var item contracts.Product
		err = mapstructure.Decode(doc.Data(), &item)
		if err != nil {
			log.Fatalln(err)
		}
		item.Id = doc.Ref.ID
		array = append(array, item)
	}

	return array
}

func (repo *ProductsRepo) FindById(id string) contracts.Product {

	client := BaseRepo()
	doc, err := client.Collection(COLLECTION_PRODUCTS).Doc(id).Get(ctx)
	var item contracts.Product
	err = mapstructure.Decode(doc.Data(), &item)
	if err != nil {
		log.Fatalln(err)
	}

	return item
}

func (repo *ProductsRepo) Save(item contracts.Product) contracts.Product {

	client := BaseRepo()
	fmt.Println(item)
	doc, _, err := client.Collection(COLLECTION_PRODUCTS).Add(ctx, item)
	if err != nil {
		log.Fatalln(err)
	}
	item.Id = doc.ID

	return item
}

func (repo *ProductsRepo) Update(id string, item contracts.Product) contracts.Product {

	client := BaseRepo()
	_, err := client.Collection(COLLECTION_PRODUCTS).Doc(id).Set(ctx, item)
	if err != nil {
		log.Fatalln(err)
	}
	item.Id = id

	return item
}

func (repo *ProductsRepo) Delete(id string) bool{

	client := BaseRepo()
	_, err := client.Collection(COLLECTION_PRODUCTS).Doc(id).Delete(ctx)
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}