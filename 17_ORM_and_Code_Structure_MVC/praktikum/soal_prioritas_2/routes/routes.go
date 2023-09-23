package routes

import (
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_2/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, usercontroller controller.UserController) {
	e.GET("/users", usercontroller.GetUsers)
	e.GET("/users/:id", usercontroller.GetUser)
	e.POST("/users", usercontroller.Create)
	e.PUT("/users/:id", usercontroller.Update)
	e.DELETE("/users/:id", usercontroller.Delete)
}

func RouteBook(e *echo.Echo, bookController controller.BookController) {
	e.GET("/books", bookController.GetBooks)
	e.GET("/books/:id", bookController.GetBook)
	e.POST("/books", bookController.Insert)
	e.PUT("/books/:id", bookController.Update)
	e.DELETE("/books/:id", bookController.Delete)
}
