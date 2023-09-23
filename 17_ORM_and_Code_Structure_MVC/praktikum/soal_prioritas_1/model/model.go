package model

import (
	"fmt"
	"go_bagus-ario-yudanto/17_ORM_and_Code_Structure_MVC/praktikum/soal_prioritas_1/configs"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel(config configs.Config) *gorm.DB {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Model : cannot connect to database, ", err.Error())
		return nil
	}

	return DB
}

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&User{})
}
