package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/burntcarrot/pm/errors"
	"github.com/burntcarrot/pm/helpers"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Usecase user.Usecase
}

func NewAuthController(u user.Usecase) *AuthController {
	return &AuthController{
		Usecase: u,
	}
}

func (a *AuthController) Login(c echo.Context) error {
	userRequest := LoginRequest{}
	err := c.Bind(&userRequest)
	if err != nil {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrBadRequest)
	}

	ctx := c.Request().Context()

	u, err := a.Usecase.Login(ctx, strings.ToLower(userRequest.Username), userRequest.Password)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusUnauthorized, errors.ErrValidationFailed)
	}
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	token, err := helpers.GenerateToken(u.Username, u.Role)
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	return controllers.Success(c, LoginResponse{Token: token})
}

func (a *AuthController) Register(c echo.Context) error {
	userRequest := RegisterRequest{}
	err := c.Bind(&userRequest)
	if err != nil {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrBadRequest)
	}

	// fetch context
	ctx := c.Request().Context()

	// TODO: check if user already exists

	// map user
	userDomain := user.Domain{
		Username: strings.ToLower(userRequest.Username),
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Role:     userRequest.Role,
	}

	// register user
	u, err := a.Usecase.Register(ctx, userDomain)
	if err == errors.ErrValidationFailed {
		return controllers.Error(c, http.StatusUnauthorized, errors.ErrValidationFailed)
	}
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	response := RegisterResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}

	fmt.Println("Woohoo register!")

	return controllers.Success(c, response)
}
