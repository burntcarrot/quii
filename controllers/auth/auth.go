package auth

import (
	"fmt"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/user"
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
	userLogin := LoginRequest{}
	err := c.Bind(&userLogin)
	if err != nil {
		return err
	}

	ctx := c.Request().Context()

	u, err := a.Usecase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		return err
	}
	fmt.Println(u)

	// generate token
	// SOLVED: inspect why user.ID is 0 => Redis doesn't use gorm, so no ID is generated
	token, err := helpers.GenerateToken(int(u.ID), u.Role)
	fmt.Println(err)
	fmt.Println(token)
	if err != nil {
		return err
	}

	return controllers.Success(c, LoginResponse{Token: token})
}

func (a *AuthController) Register(c echo.Context) error {
	userRegister := RegisterRequest{}
	err := c.Bind(&userRegister)
	if err != nil {
		return err
	}

	// fetch context
	ctx := c.Request().Context()

	// TODO: check if user already exists

	// map user
	userDomain := user.Domain{
		Email:    userRegister.Email,
		Password: userRegister.Password,
		Role:     userRegister.Role,
	}

	// register user
	u, err := a.Usecase.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	registerResponse := RegisterResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  u.Role,
	}

	fmt.Println("Woohoo register!")

	return controllers.Success(c, registerResponse)
}
