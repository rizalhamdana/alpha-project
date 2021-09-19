package repository

import (
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/health/model"
	"rizalhamdana/alpha_project/store_service/modules/share"

	"github.com/sirupsen/logrus"
)

type HealthRepository struct {
}

func NewHealthRepository() I_HealthRepository {
	return &HealthRepository{}
}

func (r *HealthRepository) GetHealthStatus(result chan<- share.ResultRepository) {
	ctx := "repo_health.get_health_status"
	resultRepo := share.ResultRepository{
		Error: nil,
		Result: &model.Health{
			Version:     "v1",
			Status:      "Running",
			ServiceName: "Store Service",
		},
	}
	helper.Log(logrus.InfoLevel, "Get Health Status Called", ctx, "")
	result <- resultRepo
}
