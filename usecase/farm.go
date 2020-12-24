package usecase

import (
	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/repo/farm"
)

// FarmUsecase is a usecase for farm
type FarmUsecase struct {
	farmRepo farm.IRepo
}

// Create is a create farm usecase
func (fc FarmUsecase) Create(farm entity.Farm, farmerID int) (fid int, err error) {
	return fc.farmRepo.CreateFarm(farm, farmerID)
}

// NewFarmUsecase is a factory method for farm usecase
func NewFarmUsecase(fr farm.IRepo) FarmUsecase {
	return FarmUsecase{
		farmRepo: fr,
	}
}
