package presenter

import (
	"net/http"
	"rizalhamdana/alpha_project/store_service/helper"
	healthModel "rizalhamdana/alpha_project/store_service/modules/api/v1/health/model"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/model"
	"rizalhamdana/alpha_project/store_service/modules/api/v1/store/usecase"
	"rizalhamdana/alpha_project/store_service/modules/share"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type StoreHttpHandler struct {
	usecase usecase.I_StoreUsecase
}

func NewStoreHttpHandler(storeUsecase usecase.I_StoreUsecase) *StoreHttpHandler {
	return &StoreHttpHandler{storeUsecase}
}

func (h *StoreHttpHandler) Mount(group *echo.Group) {
	group.GET("/store/check", h.CheckStoreAPI)
	group.GET("/store/:uuid", h.GetOneStore)
	group.POST("/store", h.CreateNewStore)
	group.PUT("/store/:uuid", h.UpdateOneStore)

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
	ctx := "store_handler.create_new_store"
	article := model.Store{}
	err = e.Bind(&article)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "Binding request body json to store")
		return e.JSON(http.StatusBadRequest, share.ResponseDetail{
			Message:   "Failed to process request, try again later",
			Data:      nil,
			IsSuccess: false},
		)
	}
	data, err := h.usecase.CreateStore(&article)
	if err != nil {
		helper.Log(logrus.ErrorLevel, err.Error(), ctx, "create article")
		return e.JSON(http.StatusBadRequest, share.ResponseDetail{
			Message:   "Failed to process request, try again later",
			Data:      nil,
			IsSuccess: false},
		)
	}
	return e.JSON(http.StatusCreated, share.ResponseDetail{
		Message:   "Successfully create store",
		Data:      data,
		IsSuccess: true,
	})
}

func (h *StoreHttpHandler) UpdateOneStore(e echo.Context) (err error) {
	return nil
}

func (h *StoreHttpHandler) DeleteOneStore(e echo.Context) (err error) {
	return nil
}
