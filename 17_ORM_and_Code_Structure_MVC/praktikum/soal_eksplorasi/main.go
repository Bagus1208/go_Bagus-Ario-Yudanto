// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/driver/mysql"
// 	_ "gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var (
// 	DB *gorm.DB
// )

// func init() {
// 	InitDB()
// 	InitialMigration()
// }

// type Config struct {
// 	DB_Username string
// 	DB_Password string
// 	DB_Port     string
// 	DB_Host     string
// 	DB_Name     string
// }

// func InitDB() {

// 	config := Config{
// 		DB_Username: "root",
// 		DB_Password: "root",
// 		DB_Port:     "3306",
// 		DB_Host:     "localhost",
// 		DB_Name:     "crud_go",
// 	}

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		config.DB_Username,
// 		config.DB_Password,
// 		config.DB_Host,
// 		config.DB_Port,
// 		config.DB_Name,
// 	)

// 	var err error
// 	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}
// }

// type User struct {
// 	gorm.Model
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }

// func InitialMigration() {
// 	DB.AutoMigrate(&User{})
// }

// func SetResponse(message string, data any) map[string]any {
// 	if data == nil {
// 		return map[string]any{
// 			"message": message,
// 		}
// 	}

// 	return map[string]any{
// 		"message": message,
// 		"data":    data,
// 	}
// }

// // get all users
// func GetUsersController(c echo.Context) error {
// 	var users []User

// 	if err := DB.Find(&users).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, SetResponse("success get all users", users))
// }

// // get user by id
// func GetUserController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}

// 	if id < 0 {
// 		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
// 	}

// 	var user User
// 	if err := DB.First(&user, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}
// 	return c.JSON(http.StatusOK, SetResponse("success get user", user))
// }

// // create new user
// func CreateUserController(c echo.Context) error {
// 	user := User{}
// 	c.Bind(&user)

// 	if err := DB.Save(&user).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create new user",
// 		"user":    user,
// 	})
// }

// // delete user by id
// func DeleteUserController(c echo.Context) error {
// 	// your solution here
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}

// 	if id <= 0 {
// 		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
// 	}

// 	var users []User
// 	if err = DB.Find(&users).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}

// 	if err = DB.Delete(&users, id).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, []int{})
// }

// // update user by id
// func UpdateUserController(c echo.Context) error {
// 	// your solution here
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}

// 	if id < 0 {
// 		return c.JSON(http.StatusNotFound, SetResponse("user not found", nil))
// 	}

// 	var userUpdate User
// 	if err := c.Bind(&userUpdate); err != nil {
// 		return c.JSON(http.StatusBadRequest, SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
// 	}

// 	var user User
// 	if err = DB.Model(&user).Where("id = ?", id).Updates(userUpdate).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, SetResponse(err.Error(), nil))
// 	}

// 	return c.JSON(http.StatusOK, SetResponse("success update", user))
// }

// func main() {
// 	InitDB()
// 	InitialMigration()
// 	// create a new echo instance
// 	e := echo.New()
// 	// Route / to handler function
// 	e.GET("/users", GetUsersController)
// 	e.GET("/users/:id", GetUserController)
// 	e.POST("/users", CreateUserController)
// 	e.DELETE("/users/:id", DeleteUserController)
// 	e.PUT("/users/:id", UpdateUserController)

// 	// start the server, and log if it fails
// 	e.Logger.Fatal(e.Start(":8000"))
// }

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
