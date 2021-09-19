package usecase

import "rizalhamdana/alpha_project/store_service/modules/api/v1/health/model"

type I_HealthUsecase interface {
	GetHealthStatus() (health *model.Health, err error)
}
