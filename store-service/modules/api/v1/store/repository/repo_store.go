package repository

import (
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/model"
	"rizalhamdana/alpha_project/store_service/modules/share"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ctx := "store_repository.insert_one_store"

	storeCollection := r.MongoWrite.Collection("store")
	newStore.Id = primitive.NewObjectID()
	newStore.CreatedAt = time.Now()
	newStore.UpdatedAt = time.Now()
	newStore.IsActive = true
	newStore.CreatedBy = "Initial Setup"
	newStore.UpdatedBy = "Initial Setup"

	insert, err := storeCollection.InsertOne(*r.Ctx, newStore)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Insert Store")
		result <- share.ResultRepository{Error: err}
		return
	}

	newStore.Id = insert.InsertedID.(primitive.ObjectID)
	result <- share.ResultRepository{
		Result: newStore,
		Error:  nil,
	}
}
func (r *StoreRepository) UpdateOneStore(updateStore *model.Store, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}
func (r *StoreRepository) DeleteOneStoreByUuid(uuid string, result chan<- share.ResultRepository) {
	result <- share.ResultRepository{}
}
