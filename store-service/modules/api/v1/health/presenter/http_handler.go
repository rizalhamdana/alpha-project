package presenter

import (
	"net/http"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/health/model"

	"github.com/labstack/echo/v4"
)

type HealthHttpHandler struct {
}

func NewHealthHandler() *HealthHttpHandler {
	return &HealthHttpHandler{}
}

func (h *HealthHttpHandler) Mount(e *echo.Group) {
	e.GET("/check", h.CheckHealth)

}

func (h *HealthHttpHandler) CheckHealth(e echo.Context) (err error) {
	// ctx := "handler_health.check_health"

	health := model.Health{
		Status:      "Running",
		Version:     "v1",
		ServiceName: "Store Service",
	}

	return e.JSON(http.StatusOK, health)
}
