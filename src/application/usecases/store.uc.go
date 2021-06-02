package usecases

import (
	"github.com/mtsluna/go-market/src/domain/contracts"
	"github.com/mtsluna/go-market/src/infrastructure/repositories"
)

type StoreUc struct {
	repo repositories.StoreRepo
}

func (uc *StoreUc) FindAll() [] contracts.Store{

	var array = uc.repo.FindAll()

	return array

}

func (uc *StoreUc) FindById(id string) contracts.Store {

	var store = uc.repo.FindById(id)

	return store

}

func (uc *StoreUc) Save(store contracts.Store) contracts.Store {

	store = uc.repo.Save(store)

	return store

}

func (uc *StoreUc) Update(id string, store contracts.Store) contracts.Store {

	store = uc.repo.Update(id, store)

	return store

}
