package usecase

import (
	"github.com/suradidchao/listenfield/entity"
	"github.com/suradidchao/listenfield/repo/farm/user"
)

// UserUsecase is a usecase for user
type UserUsecase struct {
	userRepo user.IRepo
}

// Create is an user usecase for creating user
func (uc UserUsecase) Create(user entity.User) (uid int, err error) {
	return uc.userRepo.Create(user)
}

// NewUserUsecase is a factory method for user usecase
func NewUserUsecase(ur user.IRepo) UserUsecase {
	return UserUsecase{
		userRepo: ur,
	}
}
