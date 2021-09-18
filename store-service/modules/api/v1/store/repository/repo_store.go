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

func (r *StoreRepository) SearchStore(searchParam helper.Filter) (result *share.ResultRepository) {
	return nil
}
func (r *StoreRepository) GetOneStoreByUuid(uuid string) (result *share.ResultRepository) {
	return nil
}
func (r *StoreRepository) InsertOneStore(newStore *model.Store) (result *share.ResultRepository) {
	return nil
}
func (r *StoreRepository) UpdateOneStore(updateStore *model.Store) (result *share.ResultRepository) {
	return nil
}
func (r *StoreRepository) DeleteOneStoreByUuid(uuid string) (result *share.ResultRepository) {
	return nil
}
