package usecases

import (
	"github.com/mtsluna/go-market/src/domain/contracts"
	"github.com/mtsluna/go-market/src/infrastructure/repositories"
)

type ProductsUc struct {
	repo repositories.ProductsRepo
}

func (uc *ProductsUc) FindAll() [] contracts.Product{

	var array = uc.repo.FindAll()

	return array

}

func (uc *ProductsUc) FindAllByStoreId(id string) [] contracts.Product{

	var array = uc.repo.FindAllByStoreId(id)

	return array

}

func (uc *ProductsUc) FindAllByTitle(title string) [] contracts.Product{

	var array = uc.repo.FindAllByTitle(title)

	return array

}

func (uc *ProductsUc) FindById(id string) contracts.Product {

	var product = uc.repo.FindById(id)

	return product

}

func (uc *ProductsUc) Save(product contracts.Product) contracts.Product {

	product = uc.repo.Save(product)

	return product

}

func (uc *ProductsUc) Update(id string, product contracts.Product) contracts.Product {

	product = uc.repo.Update(id, product)

	return product

}

func (uc *ProductsUc) Delete(id string) bool {

	deleted := uc.repo.Delete(id)

	return deleted

}

