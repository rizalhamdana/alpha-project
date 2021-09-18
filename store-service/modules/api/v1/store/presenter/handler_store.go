package presenter

import (
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
	group.GET("/store/:uuid", h.GetOneStore)
	group.POST("/store", h.CreateNewStore)
	group.PUT("/store/:uuid", h.UpdateOneStore)
	group.DELETE("/store/:uuid", h.DeleteOneStore)

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
