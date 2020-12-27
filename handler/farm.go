package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
		FarmOwner: entity.User{
			UserID: farmPayload.FarmOwnerID,
		},
	}
	createdFarm, err := f.farmUsecase.Create(farm, farmPayload.FarmOwnerID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to create Farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK", Data: createdFarm})
}

// AddWorker is a handler for adding a worker to a farm
func (f FarmHandler) AddWorker(c echo.Context) error {
	farmID, err := strconv.Atoi(c.Param("farm_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid farm id"})
	}
	var addFarmWorkerPayload AddFarmWorkerPayload
	if err := c.Bind(&addFarmWorkerPayload); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	addFarmWorkerOperationID, err := f.farmUsecase.AddWorker(farmID, addFarmWorkerPayload.WorkerID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to add worker to a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK", Data: addFarmWorkerOperationID})
}

// GetAllWorkers is a handler for getting all workers in a farm
func (f FarmHandler) GetAllWorkers(c echo.Context) error {
	farmID, err := strconv.Atoi(c.Param("farm_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid farm id"})
	}

	farmWorkerIDs, err := f.farmUsecase.GetAllWorkers(farmID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to get all workers of a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK", Data: farmWorkerIDs})
}

// NewFarmHandler is a factory method for farm handler
func NewFarmHandler(fu usecase.FarmUsecase) FarmHandler {
	return FarmHandler{
		farmUsecase: fu,
	}
}
