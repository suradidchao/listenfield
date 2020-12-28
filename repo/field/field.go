package field

import (
	"github.com/suradidchao/listenfield/entity"
)

// IRepo is an interface for field repository
type IRepo interface {
	Create(field entity.Field) (fid int, err error)
	Delete(fieldID int) (err error)
}

// Repo is a user repository for managing field in db
type Repo struct {
	fieldAdapter IAdapter
}

// Create is a method for creating field in db
func (r Repo) Create(field entity.Field) (fid int, err error) {
	return r.fieldAdapter.Create(field)
}

// Delete is a method for deleting a field from db
func (r Repo) Delete(fieldID int) (err error) {
	return r.fieldAdapter.Delete(fieldID)
}

// NewRepo is a factory method of field repository
func NewRepo(fa IAdapter) Repo {
	return Repo{
		fieldAdapter: fa,
	}
}
