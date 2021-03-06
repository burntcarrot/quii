package routes

import (
	"github.com/burntcarrot/quii/controllers/auth"
	"github.com/burntcarrot/quii/controllers/project"
	"github.com/burntcarrot/quii/controllers/task"
	"github.com/burntcarrot/quii/controllers/user"
	"github.com/burntcarrot/quii/helpers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Controllers struct {
	AuthController    *auth.AuthController
	UserController    *user.UserController
	ProjectController *project.ProjectController
	TaskController    *task.TaskController
}

func (c *Controllers) InitRoutes(e *echo.Echo) {
	// expose prometheus metrics
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()), middleware.CORS())

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

		api.GET("/profile/:userName", c.UserController.GetByName)
	}

	u := api.Group("/u/:userName")
	u.Use(helpers.UserRoleValidation)
	{
		// projects
		u.POST("/create", c.ProjectController.CreateProject)
		u.GET("/projects", c.ProjectController.GetProjects)
		u.GET("/projects/:projectName", c.ProjectController.GetProjectByName)

		// tasks
		u.POST("/projects/:projectName/tasks/new", c.TaskController.CreateTask)
		u.GET("/projects/:projectName/tasks", c.TaskController.GetTasks)
		u.GET("/projects/:projectName/tasks/:taskID", c.TaskController.GetTaskByID)
	}
}
