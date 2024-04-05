package database

import (
	"keraton-task-3/models"

	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func DbMigrate(dbParam *gorm.DB) {
	DbConnection = dbParam

	DbConnection.AutoMigrate(
		&models.Student{},
		&models.Lecturer{},
		&models.Course{},
		&models.CourseTaken{},
	)
}
