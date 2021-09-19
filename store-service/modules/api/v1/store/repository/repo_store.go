package repository

import (
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/model"
	"rizalhamdana/alpha_project/store_service/modules/share"
)

type StoreRepository struct {
	*share.Repository
}

func NewStoreRepository(repo *share.Repository) I_StoreRepository {
	return &StoreRepository{repo}
}

func (r *StoreRepository) SearchStore(filter helper.Filter, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}

func (r *StoreRepository) GetOneStoreByUuid(uuid string, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}
func (r *StoreRepository) InsertOneStore(newStore *model.Store, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}
func (r *StoreRepository) UpdateOneStore(updateStore *model.Store, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}
func (r *StoreRepository) DeleteOneStoreByUuid(uuid string, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}
