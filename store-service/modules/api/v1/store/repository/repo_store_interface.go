package repository

import (
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/model"
	"rizalhamdana/alpha_project/store_service/modules/share"
)

type I_StoreRepository interface {
	SearchStore(filter helper.Filter) (result *share.ResultRepository)
	GetOneStoreByUuid(uuid string) (result *share.ResultRepository)
	InsertOneStore(newStore *model.Store) (result *share.ResultRepository)
	UpdateOneStore(updateStore *model.Store) (result *share.ResultRepository)
	DeleteOneStoreByUuid(uuid string) (result *share.ResultRepository)
}
