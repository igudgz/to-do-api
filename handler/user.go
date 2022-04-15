package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"todo-api/model"
	"todo-api/service"
	"todo-api/util"
)

type UserHandler interface {
	AuthenticateUser(c echo.Context) error
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) AuthenticateUser(c echo.Context) error {
	requestLogin := new(model.UserLogin)

	if err := c.Bind(&requestLogin); err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Errorf("Invalid Parameters"))
	}

	user, valid := h.userService.ValidateUserCredentials(requestLogin.Email, requestLogin.Password)
	if !valid {
		return c.JSON(http.StatusUnauthorized, fmt.Errorf("user Invalid"))
	}

	token, err := util.GenerateJwtToken(user)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, fmt.Errorf("error in generate token"))
	}
	return c.JSON(http.StatusOK, model.Token{Token: token})
}
