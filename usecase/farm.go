package usecase

import (
	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/repo/farm"
	"github.com/suradidchao/listenfield/repo/farmworker"
)

// FarmUsecase is a usecase for farm
type FarmUsecase struct {
	farmRepo       farm.IRepo
	farmWorkerRepo farmworker.IRepo
}

// Create is a create farm usecase
func (fc FarmUsecase) Create(farm entity.Farm, farmerID int) (fid int, err error) {
	return fc.farmRepo.CreateFarm(farm, farmerID)
}

// AddWorker is a usecase for adding worker to farm
func (fc FarmUsecase) AddWorker(farmID int, workerID int) (fwID int, err error) {
	return fc.farmWorkerRepo.Create(farmID, workerID)
}

// GetAllWorkers is a usecase for adding worker to farm
func (fc FarmUsecase) GetAllWorkers(farmID int) (userIDs []int, err error) {
	return fc.farmWorkerRepo.GetAllByFarmID(farmID)
}

// NewFarmUsecase is a factory method for farm usecase
func NewFarmUsecase(fr farm.IRepo, fwr farmworker.IRepo) FarmUsecase {
	return FarmUsecase{
		farmRepo:       fr,
		farmWorkerRepo: fwr,
	}
}
