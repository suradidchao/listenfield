package farm

import "github.com/suradidchao/listenfield/repo"

// IRepo is an interface for farm repository
type IRepo interface {
	CreateFarm(farm repo.Farm, farmerID int) (farmID int, err error)
}

// Repo is a farm repo
type Repo struct {
	farmAdapter IAdapter
}

// CreateFarm is a method for creating farm for repo
func (r Repo) CreateFarm(farm repo.Farm, farmerID int) (farmID int, err error) {
	return r.farmAdapter.CreateFarm(farm, farmerID)
}

// NewRepo is a factory method for creating farm repo
func NewRepo(fa IAdapter) Repo {
	return Repo{
		farmAdapter: fa,
	}
}
