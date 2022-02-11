package main

import (
	"log"
	"time"

	"github.com/burntcarrot/pm/cmd/pm/routes"
	"github.com/burntcarrot/pm/controllers/auth"
	pc "github.com/burntcarrot/pm/controllers/project"
	tc "github.com/burntcarrot/pm/controllers/task"
	uc "github.com/burntcarrot/pm/controllers/user"
	projectDbRedis "github.com/burntcarrot/pm/drivers/db/project/redis"
	taskDbRedis "github.com/burntcarrot/pm/drivers/db/task/redis"
	userDbRedis "github.com/burntcarrot/pm/drivers/db/user/redis"
	"github.com/burntcarrot/pm/drivers/redis"
	"github.com/burntcarrot/pm/entity/project"
	"github.com/burntcarrot/pm/entity/task"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// get DB connection
	// TODO: parse DB connection configs from config file
	dbConfig := redis.DBConfig{
		User: "test",
	}
	Conn := dbConfig.InitDB()

	// set timeout
	// TODO: parse timeout duration from config file
	timeout := time.Duration(time.Minute * 1)

	// initialize usecase
	userUsecase := user.NewUsecase(userDbRedis.NewUserRepo(Conn), timeout)
	projectUsecase := project.NewUsecase(projectDbRedis.NewProjectRepo(Conn), timeout)
	taskUsecase := task.NewUsecase(taskDbRedis.NewTaskRepo(Conn), timeout)

	// initialize controllers
	authController := auth.NewAuthController(*userUsecase)
	userController := uc.NewUserController(*userUsecase)
	projectController := pc.NewProjectController(*projectUsecase)
	taskController := tc.NewTaskController(*taskUsecase)

	// initialize routes and start echo server
	rc := routes.Controllers{
		AuthController:    authController,
		UserController:    userController,
		ProjectController: projectController,
		TaskController:    taskController,
	}
	rc.InitRoutes(e)

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	log.Println(e.Start(":8080"))
}
