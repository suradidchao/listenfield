package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/suradidchao/listenfield/internal/passgen"
	"github.com/suradidchao/listenfield/repo/user"
)

// AuthorizeUsecase is a collection of usecases aboutu authorize
type AuthorizeUsecase struct {
	userRepo user.IRepo
	secret   string
}

// Authorize is an  usecase for authorizing user
func (auc AuthorizeUsecase) Authorize(username string, password string) (token string, err error) {
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
	claims["ownFarmId"] = []int{1826, 1827}
	claims["workingFarmIds"] = []int{1824, 1825}
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	token, err = jwtToken.SignedString([]byte(auc.secret))
	if err != nil {
		return token, err
	}

	return token, nil
}

// NewAuthorizeUsecase is a factory method for AuthorizeUsecase
func NewAuthorizeUsecase(ur user.IRepo, s string) AuthorizeUsecase {
	return AuthorizeUsecase{
		userRepo: ur,
		secret:   s,
	}
}
