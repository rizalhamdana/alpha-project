package presenter

import (
	"net/http"
	healthModel "rizalhamdana/alpha_project/store_service/modules/api/v1/health/model"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/usecase"

	"github.com/labstack/echo/v4"
)

type StoreHttpHandler struct {
	usecase *usecase.I_StoreUsecase
}

func NewStoreHttpHandler(storeUsecase usecase.I_StoreUsecase) *StoreHttpHandler {
	return &StoreHttpHandler{&storeUsecase}
}

func (h *StoreHttpHandler) Mount(group *echo.Group) {
	group.GET("/store/check", h.CheckStoreAPI)
	group.GET("/store/:uuid", h.GetOneStore)
	group.POST("/store", h.CreateNewStore)
	group.PUT("/store/:uuid", h.UpdateOneStore)
	group.DELETE("/store/:uuid", h.DeleteOneStore)

}

func (h *StoreHttpHandler) CheckStoreAPI(e echo.Context) (err error) {
	health := healthModel.Health{
		Status:      "Running",
		Version:     "v1",
		ServiceName: "Store API v1",
	}

	return e.JSON(http.StatusOK, health)
}

func (h *StoreHttpHandler) GetOneStore(e echo.Context) (err error) {
	return nil
}

func (h *StoreHttpHandler) CreateNewStore(e echo.Context) (err error) {
	return nil
}

func (h *StoreHttpHandler) UpdateOneStore(e echo.Context) (err error) {
	return nil
}

func (h *StoreHttpHandler) DeleteOneStore(e echo.Context) (err error) {
	return nil
}
