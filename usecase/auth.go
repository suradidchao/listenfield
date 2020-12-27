package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/suradidchao/listenfield/internal/passgen"
	"github.com/suradidchao/listenfield/repo/user"
)

// AuthUsecase is a collection of usecases about auth
type AuthUsecase struct {
	userRepo user.IRepo
	secret   string
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

	// Create token
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["ownedFarmIds"] = user.OwnedFarmIDs
	claims["workingFarmIds"] = user.WorkingFarmIDs
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	token, err = jwtToken.SignedString([]byte(auc.secret))
	if err != nil {
		return token, err
	}

	return token, nil
}

// NewAuthUsecase is a factory method for AuthorizeUsecase
func NewAuthUsecase(ur user.IRepo, s string) AuthUsecase {
	return AuthUsecase{
		userRepo: ur,
		secret:   s,
	}
}
