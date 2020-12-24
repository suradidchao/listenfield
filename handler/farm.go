package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/usecase"
)

// FarmHandler is a farm handler
type FarmHandler struct {
	farmUsecase usecase.FarmUsecase
}

// CreateFarm is a handler for create farm endpoint
func (f FarmHandler) CreateFarm(c echo.Context) error {
	var farmPayload CreateFarmPayload
	if err := c.Bind(&farmPayload); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	farm := entity.Farm{
		FarmName: farmPayload.FarmName,
		FarmOwner: entity.Farmer{
			FarmerID: farmPayload.FarmOwnerID,
		},
	}
	createdFarm, err := f.farmUsecase.Create(farm, farmPayload.FarmOwnerID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to create Farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK", Data: createdFarm})
}

// NewFarmHandler is a factory method for farm handler
func NewFarmHandler(fu usecase.FarmUsecase) FarmHandler {
	return FarmHandler{
		farmUsecase: fu,
	}
}
