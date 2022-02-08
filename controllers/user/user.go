package user

import (
	"net/http"
	"strings"

	"github.com/burntcarrot/pm/controllers"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/burntcarrot/pm/errors"
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

func (u *UserController) GetByName(c echo.Context) error {
	ctx := c.Request().Context()
	userName := c.Param("userName")
	us, err := u.Usecase.GetByName(ctx, strings.ToLower(userName))
	if err != nil {
		return controllers.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError)
	}

	response := UserResponse{
		Username: us.Username,
		Email:    us.Email,
		Role:     us.Role,
	}

	return controllers.Success(c, response)
}
