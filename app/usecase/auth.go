package usecase

import (
	"errors"

	"github.com/suradidchao/listenfield/app/lib/jwtgen"
	"github.com/suradidchao/listenfield/app/lib/passgen"
	"github.com/suradidchao/listenfield/app/repo/user"
)

// AuthUsecase is a collection of usecases about auth
type AuthUsecase struct {
	userRepo     user.IRepo
	jwtGenerator jwtgen.IJWTGenerator
}

// Authenticate is an  usecase for authorizing user
func (auc AuthUsecase) Authenticate(username string, password string) (token string, err error) {
	user, err := auc.userRepo.GetByUsername(username)
	if err != nil {
		return token, err
	}

	// Throws unauthorized error
	if username != user.Username || !passgen.ComparePasswords(user.Password, []byte(password)) {
		return token, errors.New("Unauthorized access")
	}

	claim := jwtgen.NewCustomClaim(user.Username, user.OwnedFarmIDs, user.WorkingFarmIDs)
	token, err = auc.jwtGenerator.Gen(claim)
	if err != nil {
		return token, err
	}

	return token, nil
}

// NewAuthUsecase is a factory method for AuthorizeUsecase
func NewAuthUsecase(ur user.IRepo, jwtGen jwtgen.IJWTGenerator) AuthUsecase {
	return AuthUsecase{
		userRepo:     ur,
		jwtGenerator: jwtGen,
	}
}
