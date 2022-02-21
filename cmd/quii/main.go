package main

import (
	"log"
	"time"

	"github.com/burntcarrot/quii/cmd/quii/routes"
	"github.com/burntcarrot/quii/controllers/auth"
	pc "github.com/burntcarrot/quii/controllers/project"
	tc "github.com/burntcarrot/quii/controllers/task"
	uc "github.com/burntcarrot/quii/controllers/user"
	projectDbRedis "github.com/burntcarrot/quii/drivers/db/project/redis"
	taskDbRedis "github.com/burntcarrot/quii/drivers/db/task/redis"
	userDbRedis "github.com/burntcarrot/quii/drivers/db/user/redis"
	"github.com/burntcarrot/quii/drivers/redis"
	"github.com/burntcarrot/quii/entity/project"
	"github.com/burntcarrot/quii/entity/task"
	"github.com/burntcarrot/quii/entity/user"
	"github.com/burntcarrot/quii/metrics"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	// register prometheus metrics
	_ = prometheus.Register(metrics.PromLoginRequests)
	_ = prometheus.Register(metrics.PromLoginDurations)
	_ = prometheus.Register(metrics.PromLoginRequestSizes)
}

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
	log.Println(e.Start(":8080"))
}
