package main

import (
	"clean_architecture/config"
	"clean_architecture/features/users/handler"
	"clean_architecture/features/users/repository"
	"clean_architecture/features/users/service"
	"clean_architecture/middleware"
	"clean_architecture/routes"
	"clean_architecture/utils/database"
	"fmt"

	"github.com/labstack/echo/v4"
)

var ()

func main() {
	app := echo.New()
	var userConfig = config.InitConfig()
	var db = database.ConnectDB(userConfig)
	database.Migrate(db)

	var (
		userRepository = repository.New(db)
		userservice    = service.New(userRepository)
		userController = handler.New(userservice, *userConfig)
	)

	middleware.LogMiddlewares(app)

	routes.UserRoutes(app, userController, *userConfig)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", userConfig.Server_Port)).Error())
}
