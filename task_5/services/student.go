package services

import (
	"keraton_task_5/models"

	"gorm.io/gorm"
)

func GetAllStudent(db *gorm.DB) ([]models.Student, error) {
	var students []models.Student

	result := db.Find(&students)
	if result.Error != nil {
		return nil, result.Error
	}

	return students, nil
}

func GetStudentByID(db *gorm.DB, id int) (models.Student, error) {
	var student models.Student

	result := db.First(&student, id)
	if result.Error != nil {
		return models.Student{}, result.Error
	}

	return student, nil
}
