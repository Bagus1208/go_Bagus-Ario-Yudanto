package main

import (
	"REST_API_testing/configs"
	"REST_API_testing/controller"
	"REST_API_testing/helper"
	"REST_API_testing/model"
	"REST_API_testing/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()
	var config = configs.InitConfig()

	var DB = model.InitModel(*config)
	model.Migrate(DB)

	userModel := model.NewUserModel(DB)
	userControll := controller.NewUserController(userModel, *config)

	bookModel := model.NewBookModel(DB)
	bookControll := controller.NewBookController(bookModel)

	blogModel := model.NewBlogModel(DB)
	blogControll := controller.NewBlogController(blogModel)

	helper.LogMiddlewares(e)

	routes.RouteUser(e, userControll, *config)
	routes.RouteBook(e, bookControll, *config)
	routes.RouteBlog(e, blogControll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Server_Port)).Error())
}
