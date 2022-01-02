package usecase

import (
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/model"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/repository"
	"rizalhamdana/alpha_project/store_service/modules/share"

	"github.com/sirupsen/logrus"
)

type StoreUsecase struct {
	repository repository.I_StoreRepository
}

func NewStoreUsecase(storeRepo repository.I_StoreRepository) I_StoreUsecase {
	return &StoreUsecase{storeRepo}
}

func (u *StoreUsecase) SearchStore(filter helper.Filter) (storeList *[]model.Store, err error) {
	return nil, nil
}
func (u *StoreUsecase) FindOneStore(uuid string) (store *model.Store, err error) {
	return nil, nil
}
func (u *StoreUsecase) CreateStore(newStore *model.Store) (createdStore *model.Store, err error) {
	ctx := "store_usecase.create_store"

	createChannel := make(chan share.ResultRepository)

	defer close(createChannel)

	go u.repository.InsertOneStore(newStore, createChannel)

	result := <-createChannel
	if result.Error != nil {
		helper.Log(logrus.ErrorLevel, result.Error.Error(), ctx, "Insert Article")
		return nil, result.Error
	}

	createdStore = result.Result.(*model.Store)
	err = nil
	return
}
func (u *StoreUsecase) UpdateStore(updateStore *model.Store) (updatedStore *model.Store, err error) {
	return nil, nil
}
func (u *StoreUsecase) DeleteOneStore(uuid string) (deletedStore *model.Store, err error) {
	return nil, nil
}
