package usecase

import (
	"github.com/suradidchao/listenfield/entity"
)

// Farm is a usecase for farm
type Farm struct {
	farmRepo entity.Farm
}

// Create is a create farm usecase
func (f Farm) Create(farm Farm, farmer entity.Farmer) (fid int, err error) {
	return
}
