package user

import (
	"strings"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Usecase user.Usecase
}

func NewUserController(u user.Usecase) *UserController {
	return &UserController{
		Usecase: u,
	}
}

func (u *UserController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Param("userID")
	us, err := u.Usecase.GetByID(ctx, strings.ToLower(userID))
	if err != nil {
		return err
	}

	response := UserResponse{
		Username: us.Username,
		Email:    us.Email,
		Role:     us.Role,
	}

	return controllers.Success(c, response)
}
