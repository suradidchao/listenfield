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

// DeleteWorker is a handler for deleting a worker from a farm
func (f FarmHandler) DeleteWorker(c echo.Context) error {
	farmID, err := strconv.Atoi(c.Param("farm_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid farm id"})
	}

	farmworkerID, err := strconv.Atoi(c.Param("farmworker_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid farmworker id"})
	}

	err = f.farmUsecase.DeleteWorker(farmID, farmworkerID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to delete a worker from a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK"})
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

// AddTractor is a handler for adding a tractor to a farm
func (f FarmHandler) AddTractor(c echo.Context) error {
	farmID, err := strconv.Atoi(c.Param("farm_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid farm id"})
	}
	var addTractorPayload AddTractorPayload
	if err := c.Bind(&addTractorPayload); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	tractor := entity.Tractor{
		TractorName: addTractorPayload.TractorName,
		FarmID:      farmID,
	}

	tractorID, err := f.farmUsecase.AddTractor(tractor)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to add tractor to a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK", Data: tractorID})
}

// DeleteTractor is a handler for adding a tractor to a farm
func (f FarmHandler) DeleteTractor(c echo.Context) error {
	tractorID, err := strconv.Atoi(c.Param("tractor_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid tractor id"})
	}

	err = f.farmUsecase.DeleteTractor(tractorID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to delete tractor from a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK"})
}

// UpdateTractor is a handler for adding a tractor to a farm
func (f FarmHandler) UpdateTractor(c echo.Context) error {
	tractorID, err := strconv.Atoi(c.Param("tractor_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid tractor id"})
	}
	var updateTractorPayload UpdateTractorPayload
	if err := c.Bind(&updateTractorPayload); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	tractor := entity.Tractor{
		TractorName: updateTractorPayload.TractorName,
		FarmID:      updateTractorPayload.FarmID,
	}

	err = f.farmUsecase.UpdateTractor(tractorID, tractor)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to update tractor"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK"})
}

// AddField is a handler for adding a tractor to a farm
func (f FarmHandler) AddField(c echo.Context) error {
	farmID, err := strconv.Atoi(c.Param("farm_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid farm id"})
	}
	var addFieldPayload AddFieldPayload
	if err := c.Bind(&addFieldPayload); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	field := entity.Field{
		FieldName: addFieldPayload.FieldName,
		FarmID:    farmID,
		Crop:      addFieldPayload.Crop,
		Area:      addFieldPayload.Area,
	}

	fieldID, err := f.farmUsecase.AddField(field)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to add field to a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK", Data: fieldID})
}

// DeleteField is a handler for deleting a field from a farm
func (f FarmHandler) DeleteField(c echo.Context) error {
	fieldID, err := strconv.Atoi(c.Param("field_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid field id"})
	}

	err = f.farmUsecase.DeleteField(fieldID)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to delete field from a farm"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK"})
}

// UpdateField is a handler for adding a tractor to a farm
func (f FarmHandler) UpdateField(c echo.Context) error {
	fieldID, err := strconv.Atoi(c.Param("field_id"))
	if err != nil {
		fmt.Printf("Err: %s", err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid field id"})
	}
	var updateFieldPayload UpdateFieldPayload
	if err := c.Bind(&updateFieldPayload); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	field := entity.Field{
		FieldName: updateFieldPayload.FieldName,
		FarmID:    updateFieldPayload.FarmID,
		Crop:      updateFieldPayload.Crop,
		Area:      updateFieldPayload.Area,
	}

	err = f.farmUsecase.UpdateField(fieldID, field)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to update field"})
	}
	return c.JSON(http.StatusOK, Response{Message: "OK"})
}

// NewFarmHandler is a factory method for farm handler
func NewFarmHandler(fu usecase.FarmUsecase) FarmHandler {
	return FarmHandler{
		farmUsecase: fu,
	}
}
