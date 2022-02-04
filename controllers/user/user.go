package user

import (
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

	// TODO: probably use a Redis nested list sort of thing
	// userID:
	// - details
	// userID:profile
	// 		- profile details
	// userID:projects
	// 		- projects: JSON raw list
	idParam := c.Param("userID")
	us, err := u.Usecase.GetByID(ctx, idParam)
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
