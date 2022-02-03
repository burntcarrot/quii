package main

import (
	"log"
	"time"

	"github.com/burntcarrot/pm/cmd/pm/routes"
	"github.com/burntcarrot/pm/controllers/auth"
	uc "github.com/burntcarrot/pm/controllers/user"
	userDbRedis "github.com/burntcarrot/pm/drivers/db/user/redis"
	"github.com/burntcarrot/pm/drivers/redis"
	"github.com/burntcarrot/pm/entity/user"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbConfig := redis.DBConfig{
		User: "test",
	}
	Conn := dbConfig.InitDB()

	userUsecase := user.NewUsecase(userDbRedis.NewUserRepo(Conn), 60*time.Second)
	authController := auth.NewAuthController(*userUsecase)
	userController := uc.NewUserController(*userUsecase)

	rc := routes.Controllers{
		AuthController: authController,
		UserController: userController,
	}
	rc.InitRoutes(e)
	log.Println(e.Start(":8080"))
}
