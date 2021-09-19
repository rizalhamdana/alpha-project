package repository

import (
	"rizalhamdana/alpha_project/store_service/modules/share"
)

type I_HealthRepository interface {
	GetHealthStatus(result chan<- share.ResultRepository)
}
