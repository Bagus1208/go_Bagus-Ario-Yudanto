package routes

import (
	"clean_architecture/config"
	"clean_architecture/controller"
	"clean_architecture/usecase"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userusecase usecase.UserUseCase, usercontroller controller.UserController, config config.Config) {
	e.POST("/users", usercontroller.CreateUser())
	e.POST("/users/login", usercontroller.LoginUser())
	e.GET("/users", usercontroller.GetAllUsers(), echojwt.JWT([]byte(config.Secret)))
}
