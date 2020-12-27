package user

import "github.com/suradidchao/listenfield/entity"

// IRepo is an interface for user repository
type IRepo interface {
	Create(user entity.User) (uid int, err error)
	GetByUsername(username string) (user entity.User, err error)
}

// Repo is a user repository for managing user in db
type Repo struct {
	userAdapter IAdapter
}

// Create is a method for creating user in db
func (r Repo) Create(user entity.User) (uid int, err error) {
	return r.userAdapter.Create(user)
}

// GetByUsername is a method for getting a user from db by username
func (r Repo) GetByUsername(username string) (user entity.User, err error) {
	return r.userAdapter.GetByUsername(username)
}

// NewRepo is a factory method of user repository
func NewRepo(us IAdapter) Repo {
	return Repo{
		userAdapter: us,
	}
}
