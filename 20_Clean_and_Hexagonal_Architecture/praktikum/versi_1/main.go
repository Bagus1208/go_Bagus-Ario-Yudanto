package main

import (
	"clean_architecture/config"
	"clean_architecture/controller"
	"clean_architecture/repository"
	"clean_architecture/routes"
	"clean_architecture/usecase"
	"fmt"

	"github.com/labstack/echo/v4"
)

var ()

func main() {
	var userConfig = config.InitConfig()
	var db = config.ConnectDB(userConfig)

	app := echo.New()

	var (
		userRepository = repository.NewUserRepository(db)
		userUseCase    = usecase.NewUserUseCase(userRepository)
		userController = controller.NewUserController(userUseCase, *userConfig)
	)

	routes.UserRoutes(app, userUseCase, userController, *userConfig)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", userConfig.Server_Port)).Error())
}
