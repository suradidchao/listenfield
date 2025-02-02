package usecase

import (
	"time"

	"github.com/suradidchao/listenfield/app/entity"
	"github.com/suradidchao/listenfield/app/repo/activity"
	"github.com/suradidchao/listenfield/app/repo/farm"
	"github.com/suradidchao/listenfield/app/repo/farmworker"
	"github.com/suradidchao/listenfield/app/repo/field"
	"github.com/suradidchao/listenfield/app/repo/tractor"
)

// FarmUsecase is a usecase for farm
type FarmUsecase struct {
	farmRepo       farm.IRepo
	farmWorkerRepo farmworker.IRepo
	tractorRepo    tractor.IRepo
	fieldRepo      field.IRepo
	activityRepo   activity.IRepo
}

// Create is a create farm usecase
func (fc FarmUsecase) Create(farm entity.Farm, farmerID int) (fid int, err error) {
	return fc.farmRepo.CreateFarm(farm, farmerID)
}

// AddWorker is a usecase for adding worker to farm
func (fc FarmUsecase) AddWorker(farmID int, workerID int) (fwID int, err error) {
	return fc.farmWorkerRepo.Create(farmID, workerID)
}

// DeleteWorker is a usecase for adding worker to farm
func (fc FarmUsecase) DeleteWorker(farmID int, workerID int) (err error) {
	return fc.farmWorkerRepo.Delete(farmID, workerID)
}

// GetAllWorkers is a usecase for adding worker to farm
func (fc FarmUsecase) GetAllWorkers(farmID int) (userIDs []int, err error) {
	return fc.farmWorkerRepo.GetAllByFarmID(farmID)
}

// AddTractor is a usecase for adding tractor to farm
func (fc FarmUsecase) AddTractor(tractor entity.Tractor) (fwID int, err error) {
	return fc.tractorRepo.Create(tractor)
}

// DeleteTractor is a usecase for deleting tractor from a farm
func (fc FarmUsecase) DeleteTractor(tractorID int) (err error) {
	return fc.tractorRepo.Delete(tractorID)
}

// UpdateTractor is a usecase for updating a tractor's attribute in a farm
func (fc FarmUsecase) UpdateTractor(tractorID int, tractor entity.Tractor) (err error) {
	return fc.tractorRepo.Update(tractorID, tractor)
}

// AddField is a usecase for adding field to farm
func (fc FarmUsecase) AddField(field entity.Field) (fID int, err error) {
	return fc.fieldRepo.Create(field)
}

// DeleteField is a usecase for deleting a field from a farm
func (fc FarmUsecase) DeleteField(fieldID int) (err error) {
	return fc.fieldRepo.Delete(fieldID)
}

// UpdateField is a usecase for updating a field in a farm
func (fc FarmUsecase) UpdateField(fieldID int, field entity.Field) (err error) {
	return fc.fieldRepo.Update(fieldID, field)
}

// AddActivity is a usecase for adding activity to farm
func (fc FarmUsecase) AddActivity(activity entity.Activity) (aID int, err error) {
	const (
		PREP       = 10.00
		SOWED      = 20.00
		FERTILIZED = 30.00
		HARVESTED  = 100.00
	)
	var cost, revenue float64
	switch activity.ActivityName {
	case "prep":
		cost = cost + activity.Area*PREP
	case "sowed":
		cost = cost + activity.Area*SOWED
	case "fertilized":
		cost = cost + activity.Area*FERTILIZED
	case "harvested":
		revenue = revenue + activity.Area*HARVESTED
	}
	activity.Cost = cost
	activity.Revenue = revenue
	return fc.activityRepo.Create(activity)
}

// GetCostSummary is a usecase for getting cost summary of a farm within specified time period
func (fc FarmUsecase) GetCostSummary(farmID int, startDate time.Time, endDate time.Time) (cs entity.CostSummary, err error) {
	costAndRevenue, err := fc.activityRepo.GetCostRevenue(farmID, startDate, endDate)
	if err != nil {
		return cs, err
	}
	cs = entity.CostSummary{
		Cost:    costAndRevenue.Cost,
		Revenue: costAndRevenue.Revenue,
		Profit:  costAndRevenue.Revenue - costAndRevenue.Cost,
	}
	return cs, nil
}

// NewFarmUsecase is a factory method for farm usecase
func NewFarmUsecase(fr farm.IRepo, fwr farmworker.IRepo, tr tractor.IRepo, fdr field.IRepo, ar activity.IRepo) FarmUsecase {
	return FarmUsecase{
		farmRepo:       fr,
		farmWorkerRepo: fwr,
		tractorRepo:    tr,
		fieldRepo:      fdr,
		activityRepo:   ar,
	}
}
