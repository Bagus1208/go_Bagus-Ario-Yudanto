package routes

import (
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, usercontroller controller.UserController) {
	e.GET("/users", usercontroller.GetUsersController)
	e.GET("/users/:id", usercontroller.GetUserController)
	e.POST("/users", usercontroller.CreateUserController)
	e.PUT("/users/:id", usercontroller.UpdateUserController)
	e.DELETE("/users/:id", usercontroller.DeleteUserController)
}
