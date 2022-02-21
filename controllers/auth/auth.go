package auth

import (
	"net/http"
	"strings"
	"unsafe"

	"github.com/burntcarrot/quii/controllers"
	"github.com/burntcarrot/quii/entity/user"
	"github.com/burntcarrot/quii/errors"
	"github.com/burntcarrot/quii/helpers"
	"github.com/burntcarrot/quii/metrics"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
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
	// start timer
	timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
		us := v * 1000000
		metrics.PromLoginDurations.Observe(us)
	}))
	defer timer.ObserveDuration()

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

	// only increment the counter when request is successful
	metrics.PromLoginRequests.Inc()
	defer metrics.PromLoginRequestSizes.Observe(float64(unsafe.Sizeof(LoginResponse{Token: token})))

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

	us, _ := a.Usecase.GetByName(ctx, userRequest.Username)
	if us.Username != "" {
		return controllers.Error(c, http.StatusBadRequest, errors.ErrUserAlreadyExists)
	}

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

	return controllers.Success(c, response)
}
