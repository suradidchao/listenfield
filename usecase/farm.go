package usecase

import (
	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/repo/farm"
)

// Farm is a usecase for farm
type Farm struct {
	farmRepo farm.IRepo
}

// Create is a create farm usecase
func (f Farm) Create(farm entity.Farm, farmerID int) (fid int, err error) {
	return f.farmRepo.CreateFarm(farm, farmerID)
}

// NewFarm is a factory method for farm usecase
func NewFarm(fr farm.IRepo) Farm {
	return Farm{
		farmRepo: fr,
	}
}
