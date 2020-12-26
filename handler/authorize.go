package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// AuthorizeHandler is a handler for authorize
type AuthorizeHandler struct{}

// Authorize is a handler for granting jwt
func (a AuthorizeHandler) Authorize(c echo.Context) error {
	var authorizePayload AuthorizePayload
	err := c.Bind(&authorizePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	// Throws unauthorized error
	if authorizePayload.Username != "suradid" || authorizePayload.Password != "chao" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// JWT claims
	// {
	// 	"username": "suradid.c",
	// 	"ownFarmId": 1234,
	// 	"workingFarmIds": [],
	//	"exp": 1235123
	// }

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = "suradid chao"
	claims["ownFarmId"] = 0
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
func NewAuthorizeHandler() AuthorizeHandler {
	return AuthorizeHandler{}
}
