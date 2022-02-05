package routes

import (
	"github.com/burntcarrot/pm/controllers/auth"
	"github.com/burntcarrot/pm/controllers/project"
	"github.com/burntcarrot/pm/controllers/user"
	"github.com/burntcarrot/pm/helpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Controllers struct {
	AuthController    *auth.AuthController
	UserController    *user.UserController
	ProjectController *project.ProjectController
}

func (c *Controllers) InitRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// recovers from panics
	api.Use(middleware.Recover())
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${status} ${time_rfc3339} ${latency_human}\n",
	}))

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// unprotected routes
	{
		api.POST("/register", c.AuthController.Register)
		api.POST("/login", c.AuthController.Login)

		api.GET("/u/:userID", c.UserController.GetByID)
	}

	u := api.Group("/u/:userID")
	u.Use(helpers.UserRoleValidation)
	{
		u.POST("/create", c.ProjectController.CreateProject)
		u.GET("/projects", c.ProjectController.GetProjects)
	}
}
