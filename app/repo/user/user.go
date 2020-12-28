package user

import (
	"github.com/suradidchao/listenfield/app/entity"
	"github.com/suradidchao/listenfield/app/repo/farm"
	"github.com/suradidchao/listenfield/app/repo/farmworker"
)

// IRepo is an interface for user repository
type IRepo interface {
	Create(user entity.User) (uid int, err error)
	GetByUsername(username string) (user entity.User, err error)
}

// Repo is a user repository for managing user in db
type Repo struct {
	userAdapter       IAdapter
	farmAdapter       farm.IAdapter
	farmWorkerAdapter farmworker.IAdapter
}

// Create is a method for creating user in db
func (r Repo) Create(user entity.User) (uid int, err error) {
	return r.userAdapter.Create(user)
}

// GetByUsername is a method for getting a user from db by username
func (r Repo) GetByUsername(username string) (user entity.User, err error) {
	user, err = r.userAdapter.GetByUsername(username)
	if err != nil {
		return user, err
	}

	ownedFarmIDs, err := r.farmAdapter.GetFarmIDsByUserID(user.UserID)
	if err != nil {
		return user, err
	}
	user.OwnedFarmIDs = ownedFarmIDs

	workingFarmIDs, err := r.farmWorkerAdapter.GetFarmIDsByUserID(user.UserID)
	if err != nil {
		return user, err
	}
	user.WorkingFarmIDs = workingFarmIDs
	return user, nil
}

// NewRepo is a factory method of user repository
func NewRepo(ua IAdapter, fa farm.IAdapter, fwa farmworker.IAdapter) Repo {
	return Repo{
		userAdapter:       ua,
		farmAdapter:       fa,
		farmWorkerAdapter: fwa,
	}
}
