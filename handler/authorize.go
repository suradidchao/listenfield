package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/usecase"
)

// AuthorizeHandler is a handler for authorize
type AuthorizeHandler struct {
	authorizeUsecase usecase.AuthorizeUsecase
}

// Authorize is a handler for granting jwt
func (ah AuthorizeHandler) Authorize(c echo.Context) error {
	var authorizePayload AuthorizePayload
	err := c.Bind(&authorizePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}

	jwt, err := ah.authorizeUsecase.Authorize(authorizePayload.Username, authorizePayload.Password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": jwt,
	})
}

// NewAuthorizeHandler is a factory method for authorize handler
func NewAuthorizeHandler(auc usecase.AuthorizeUsecase) AuthorizeHandler {
	return AuthorizeHandler{
		authorizeUsecase: auc,
	}
}
