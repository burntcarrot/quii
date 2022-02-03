package routes

import (
	"github.com/burntcarrot/pm/controllers/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controllers struct {
	AuthController *auth.AuthController
}

func (c *Controllers) InitRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// recovers from panics
	api.Use(middleware.Recover())
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	// unprotected routes
	{
		api.POST("/register", c.AuthController.Register)
	}
}
