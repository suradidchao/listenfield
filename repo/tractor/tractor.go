package tractor

import "github.com/suradidchao/listenfield/entity"

// IRepo is an interface for tractor repository
type IRepo interface {
	Create(tractor entity.Tractor) (tid int, err error)
	Delete(tractorID int) (err error)
	Update(tractorID int, tractor entity.Tractor) (err error)
}

// Repo is a tractor repository for managing tractorr in db
type Repo struct {
	tractorAdapter IAdapter
}

// Create is a method for creating tractor in db
func (r Repo) Create(tractor entity.Tractor) (tid int, err error) {
	return r.tractorAdapter.Create(tractor)
}

// Delete is a method for deleting tractor in db
func (r Repo) Delete(tractorID int) (err error) {
	return r.tractorAdapter.Delete(tractorID)
}

// Update is a method for updating tractor attributes in a db
func (r Repo) Update(tractorID int, tractor entity.Tractor) (err error) {
	return r.tractorAdapter.Update(tractorID, tractor)
}

// NewRepo is a factory method of tractor repository
func NewRepo(ta IAdapter) Repo {
	return Repo{
		tractorAdapter: ta,
	}
}
