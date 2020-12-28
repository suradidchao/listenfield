package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suradidchao/listenfield/app/entity"
	"github.com/suradidchao/listenfield/app/usecase"
)

// UserHandler is a handler for user
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// Create is a handler for user create endpoint
func (uh UserHandler) Create(c echo.Context) error {
	var userCreatePayload UserCreatePayload
	err := c.Bind(&userCreatePayload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Invalid payload in the requests"})
	}
	user := entity.User{
		Username: userCreatePayload.Username,
		Password: userCreatePayload.Password,
		Email:    userCreatePayload.Email,
	}
	newUserID, err := uh.userUsecase.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, Response{Message: "User created", Data: newUserID})
}

// NewUserHandler is a factory methof for user handler
func NewUserHandler(uuc usecase.UserUsecase) UserHandler {
	return UserHandler{
		userUsecase: uuc,
	}
}
