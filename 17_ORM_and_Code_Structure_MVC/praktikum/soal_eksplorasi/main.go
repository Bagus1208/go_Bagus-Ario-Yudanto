package main

import (
	"fmt"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/configs"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/controller"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/model"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_eksplorasi/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()
	var config = configs.InitConfig()

	var DB = model.InitModel(*config)
	model.Migrate(DB)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel)

	bookModel := model.BookModel{}
	bookModel.Init(DB)

	bookControll := controller.BookController{}
	bookControll.InitBookController(bookModel)

	blogModel := model.BlogModel{}
	blogModel.Init(DB)

	blogControll := controller.BlogController{}
	blogControll.InitBlogController(blogModel)

	routes.RouteUser(e, userControll)
	routes.RouteBook(e, bookControll)
	routes.RouteBlog(e, blogControll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Server_Port)).Error())
}
