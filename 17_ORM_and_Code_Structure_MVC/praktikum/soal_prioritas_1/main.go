package main

import (
	"fmt"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/configs"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/controller"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/model"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()
	var config = configs.InitConfig()
	var DB = model.InitModel(*config)
	model.Migrate(DB)

	usermodel := model.UserModel{}
	usermodel.Init(DB)

	usercontroll := controller.UserController{}
	usercontroll.InitUserController(usermodel)

	routes.RouteUser(e, usercontroll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Server_Port)).Error())
}
