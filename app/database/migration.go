package database

import (
	//_projectData "Laptop/controllers/product/data"
	//_taskData "Laptop/controllers/task/data"
	_productData "Laptop/features/product/data"
	_userData "Laptop/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	db.AutoMigrate(&_productData.Product{})
}
