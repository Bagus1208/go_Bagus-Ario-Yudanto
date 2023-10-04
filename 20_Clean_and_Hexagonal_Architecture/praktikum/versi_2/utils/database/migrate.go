package database

import (
	"clean_architecture/features/users/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
