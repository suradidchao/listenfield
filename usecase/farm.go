package usecase

import (
	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/repo/farm"
	"github.com/suradidchao/listenfield/repo/farmworker"
	"github.com/suradidchao/listenfield/repo/field"
	"github.com/suradidchao/listenfield/repo/tractor"
)

// FarmUsecase is a usecase for farm
type FarmUsecase struct {
	farmRepo       farm.IRepo
	farmWorkerRepo farmworker.IRepo
	tractorRepo    tractor.IRepo
	fieldRepo      field.IRepo
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

// AddField is a usecase for adding field to farm
func (fc FarmUsecase) AddField(field entity.Field) (fID int, err error) {
	return fc.fieldRepo.Create(field)
}

// NewFarmUsecase is a factory method for farm usecase
func NewFarmUsecase(fr farm.IRepo, fwr farmworker.IRepo, tr tractor.IRepo, fdr field.IRepo) FarmUsecase {
	return FarmUsecase{
		farmRepo:       fr,
		farmWorkerRepo: fwr,
		tractorRepo:    tr,
		fieldRepo:      fdr,
	}
}
