package presenter

import (
	"net/http"
	"rizalhamdana/alpha_project/store_service/helper"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/health/usecase"
	"rizalhamdana/alpha_project/store_service/modules/share"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type HealthHttpHandler struct {
	usecase usecase.I_HealthUsecase
}

func NewHealthHandler(usecase usecase.I_HealthUsecase) *HealthHttpHandler {
	return &HealthHttpHandler{usecase: usecase}
}

func (h *HealthHttpHandler) Mount(e *echo.Group) {
	e.GET("/check", h.CheckHealth)

}

func (h *HealthHttpHandler) CheckHealth(e echo.Context) (err error) {
	ctx := "handler_health.check_health"

	health, err := h.usecase.GetHealthStatus()
	response := share.ResponseDetail{}
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "")
		response.Data = nil
		response.Message = err.Error()
		return e.JSON(http.StatusInternalServerError, response)
	}
	response.Message = "Success"
	response.Data = health
	return e.JSON(http.StatusOK, response)
}
