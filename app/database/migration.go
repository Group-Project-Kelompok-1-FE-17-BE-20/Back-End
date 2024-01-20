package database

import (
	//_projectData "Laptop/controllers/product/data"
	//_taskData "Laptop/controllers/task/data"
	_userData "Laptop/features/user/data"

	"gorm.io/gorm"
)

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
	//db.AutoMigrate(&_projectData.Project{})
	//db.AutoMigrate(&_taskData.Task{})
}
