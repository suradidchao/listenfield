package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/internal/passgen"
	"github.com/suradidchao/listenfield/usecase"
)

// AuthorizeHandler is a handler for authorize
type AuthorizeHandler struct {
	userUsecase usecase.UserUsecase
}

// Authorize is a handler for granting jwt
func (ah AuthorizeHandler) Authorize(c echo.Context) error {
	var authorizePayload AuthorizePayload
	err := c.Bind(&authorizePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	user, err := ah.userUsecase.GetByUsername(authorizePayload.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "No user found"})
	}

	// Throws unauthorized error
	if authorizePayload.Username != user.Username || !passgen.ComparePasswords(user.Password, []byte(authorizePayload.Password)) {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// JWT claims
	// {
	// 	"username": "suradid.c",
	// 	"ownFarmIds": [],
	// 	"workingFarmIds": [],
	//	"exp": 1235123
	// }

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["ownFarmId"] = []int{1826, 1827}
	claims["workingFarmIds"] = []int{1824, 1825}
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("listenfield secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// NewAuthorizeHandler is a factory method for authorize handler
func NewAuthorizeHandler(uuc usecase.UserUsecase) AuthorizeHandler {
	return AuthorizeHandler{
		userUsecase: uuc,
	}
}
