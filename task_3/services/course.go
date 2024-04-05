package services

import (
	"fmt"
	"keraton-task-3/models"

	"gorm.io/gorm"
)

func GetCourseTakenByStudent(db *gorm.DB, idStudent int) ([]models.CourseTaken, error) {
	var courses []models.CourseTaken

	result := db.Where("student_id = ?", idStudent).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}

func GetCourseTakenDetail(db *gorm.DB, idStudent int) ([]models.CourseTakenDetail, error) {
	var courses []models.CourseTakenDetail

	sqlQuery := fmt.Sprintf(`SELECT ct.id, ct.course_id, c.course_name, ct.number_semester, l.id as lecturer_id, l.lecturer_name, ct.student_id, ct.number_semester, ct.mid_exam_date, ct.mid_exam_score, ct.end_exam_date, ct.end_exam_score, ct.final_score
	FROM course_takens ct
	INNER JOIN courses c ON ct.course_id = c.id
	INNER JOIN lecturers l ON c.lecturer_id = l.id
	WHERE ct.student_id = '%d';`, idStudent)

	result := db.Raw(sqlQuery).Scan(&courses)
	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}
