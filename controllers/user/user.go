package user

import (
	"net/http"
	"strings"
	"unsafe"

	"github.com/burntcarrot/quii/controllers"
	"github.com/burntcarrot/quii/entity/user"
	"github.com/burntcarrot/quii/errors"
	"github.com/burntcarrot/quii/metrics"
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

	defer metrics.PromLoginRequestSizes.Observe(float64(unsafe.Sizeof(response)))

	return controllers.Success(c, response)
}
