package routes

import (
	"clean_architecture/config"
	"clean_architecture/features/users/entity"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userhandler entity.Handler, config config.Config) {
	e.POST("/users", userhandler.Create())
	e.POST("/users/login", userhandler.Login())
	e.GET("/users", userhandler.GetAll(), echojwt.JWT([]byte(config.Secret)))
}
