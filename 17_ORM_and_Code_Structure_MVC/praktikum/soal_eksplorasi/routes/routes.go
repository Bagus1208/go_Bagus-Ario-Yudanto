package routes

import (
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/controller"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, usercontroller controller.UserController) {
	e.GET("/users", usercontroller.GetUsers)
	e.GET("/users/:id", usercontroller.GetUser)
	e.POST("/users", usercontroller.Create)
	e.PUT("/users/:id", usercontroller.Update)
	e.DELETE("/users/:id", usercontroller.Delete)
	e.GET("/users/blogs", usercontroller.GetBlogs)
}

func RouteBook(e *echo.Echo, bookController controller.BookController) {
	e.GET("/books", bookController.GetBooks)
	e.GET("/books/:id", bookController.GetBook)
	e.POST("/books", bookController.Insert)
	e.PUT("/books/:id", bookController.Update)
	e.DELETE("/books/:id", bookController.Delete)
}

func RouteBlog(e *echo.Echo, blogController controller.BlogController) {
	e.GET("/blogs", blogController.GetBlogs)
	e.GET("/blogs/:id", blogController.GetBlog)
	e.POST("/blogs", blogController.Create)
	e.PUT("/blogs/:id", blogController.Update)
	e.DELETE("/blogs/:id", blogController.Delete)
}
