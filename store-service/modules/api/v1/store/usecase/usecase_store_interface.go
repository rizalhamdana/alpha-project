package usecase

import (
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/model"
)

type I_StoreUsecase interface {
	SearchStore(filter helper.Filter) (storeList *[]model.Store, err error)
	FindOneStore(uuid string) (store *model.Store, err error)
	CreateStore(newStore *model.Store) (createdStore *model.Store, err error)
	UpdateStore(updateStore *model.Store) (updatedStore *model.Store, err error)
	DeleteOneStore(uuid string) (deletedStore *model.Store, err error)
}
