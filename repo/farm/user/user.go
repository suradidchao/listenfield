package user

import "github.com/suradidchao/listenfield/entity"

// IRepo is an interface for user repository
type IRepo interface {
	Create(user entity.User) (uid int, err error)
}

// Repo is a user repository for managing user in mysql
type Repo struct {
	userAdapter IAdapter
}

// Create is a method for creating user in mysql
func (r Repo) Create(user entity.User) (uid int, err error) {
	return r.userAdapter.Create(user)
}

// NewRepo is a factory method of user repository
func NewRepo(us IAdapter) Repo {
	return Repo{
		userAdapter: us,
	}
}
