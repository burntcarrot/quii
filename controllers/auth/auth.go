package auth

import (
	"fmt"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/user"
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

func (a *AuthController) Register(c echo.Context) error {
	userRegister := RegisterRequest{}
	err := c.Bind(&userRegister)
	if err != nil {
		return err
	}

	// fetch context
	ctx := c.Request().Context()

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
		Email: u.Email,
		Role:  u.Role,
	}

	fmt.Println("Woohoo register!")

	return controllers.Success(c, registerResponse)
}
