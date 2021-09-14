package main

import (
	"context"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/health/presenter"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	DBRead       *mongo.Client
	DBWrite      *mongo.Client
	MongoContext context.Context

	HealthHandler *presenter.HealthHttpHandler
}

func MakeHandler() *Service {

	healthHandler := presenter.NewHealthHandler()
	return &Service{
		HealthHandler: healthHandler,
	}
}
