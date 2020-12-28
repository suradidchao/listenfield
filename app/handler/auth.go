package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/app/usecase"
)

// AuthHandler is a handler for authorize
type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

// Authenticate is a handler for granting jwt
func (ah AuthHandler) Authenticate(c echo.Context) error {
	var authorizePayload AuthorizePayload
	err := c.Bind(&authorizePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	jwt, err := ah.authUsecase.Authenticate(authorizePayload.Username, authorizePayload.Password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": jwt,
	})
}

// NewAuthHandler is a factory method for authorize handler
func NewAuthHandler(auc usecase.AuthUsecase) AuthHandler {
	return AuthHandler{
		authUsecase: auc,
	}
}
