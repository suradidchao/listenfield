package activity

import "github.com/suradidchao/listenfield/entity"

// IRepo is an interface for activity repository
type IRepo interface {
	Create(user entity.Activity) (aid int, err error)
}

// Repo is an activity repository for managing activity in db
type Repo struct {
	activityAdapter IAdapter
}

// Create is a method for creating activity in db
func (r Repo) Create(activity entity.Activity) (aid int, err error) {
	return r.activityAdapter.Create(activity)
}

// NewRepo is a factory method of activity repository
func NewRepo(aa IAdapter) Repo {
	return Repo{
		activityAdapter: aa,
	}
}
