package main

import (
	"context"
	"fmt"
	"rizalhamdana/alpha_project/store_service/config"

	healthPresenter "rizalhamdana/alpha_project/store_service/modules/api/v1/health/presenter"
	healthRepository "rizalhamdana/alpha_project/store_service/modules/api/v1/health/repository"
	healthUsecase "rizalhamdana/alpha_project/store_service/modules/api/v1/health/usecase"

	storePresenter "rizalhamdana/alpha_project/store_service/modules/api/v1/store/presenter"
	storeRepository "rizalhamdana/alpha_project/store_service/modules/api/v1/store/repository"
	storeUsecase "rizalhamdana/alpha_project/store_service/modules/api/v1/store/usecase"
	"rizalhamdana/alpha_project/store_service/modules/share"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DBRead       *mongo.Client
	DBWrite      *mongo.Client
	MongoContext context.Context

	HealthHandler *healthPresenter.HealthHttpHandler
	StoreHandler  *storePresenter.StoreHttpHandler
}

func MakeHandler() *Service {

	ctx := context.Background()

	dbRead, err := config.MongoRead(&ctx)
	if err != nil {
		fmt.Printf("Failed to connect to DB, msg: %s \n", err.Error())
		return nil
	}

	dbWrite, err := config.MongoWrite(&ctx)
	if err != nil {
		fmt.Printf("Failed to connect to DB, msg: %s \n", err.Error())
		return nil
	}

	baseRepository := share.NewRepository(dbRead, dbWrite, &ctx)
	healthRepo := healthRepository.NewHealthRepository()
	healthUsecase := healthUsecase.NewHealthUsecase(healthRepo)
	healthHandler := healthPresenter.NewHealthHandler(healthUsecase)

	storeRepo := storeRepository.NewStoreRepository(baseRepository)
	storeUsecase := storeUsecase.NewStoreUsecase(storeRepo)
	storeHandler := storePresenter.NewStoreHttpHandler(storeUsecase)
	return &Service{
		HealthHandler: healthHandler,
		StoreHandler:  storeHandler,
	}
}
