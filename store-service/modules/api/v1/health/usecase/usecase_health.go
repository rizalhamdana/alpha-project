package usecase

import (
	"errors"
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/health/model"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/health/repository"
	"rizalhamdana/alpha_project/store_service/modules/share"

	"github.com/sirupsen/logrus"
)

type HealthUsecase struct {
	healthRepo repository.I_HealthRepository
}

func NewHealthUsecase(repo repository.I_HealthRepository) I_HealthUsecase {
	return &HealthUsecase{healthRepo: repo}
}

func (u *HealthUsecase) GetHealthStatus() (health *model.Health, err error) {
	ctx := "usecase_health.get_health_status"

	repoOut := make(chan share.ResultRepository)

	go u.healthRepo.GetHealthStatus(repoOut)

	res := <-repoOut
	if res.Error != nil {
		helper.Log(logrus.ErrorLevel, res.Error.Error(), ctx, "")
		return nil, errors.New("service is not working")
	}
	return res.Result.(*model.Health), nil
}
