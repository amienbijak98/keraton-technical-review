package controllers

import (
	"keraton-task-3/database"
	"keraton-task-3/helpers"
	"keraton-task-3/models"
	"keraton-task-3/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllStudent(c *fiber.Ctx) error {
	students, err := services.GetAllStudent(database.DbConnection)

	if err != nil {
		return helpers.ErrorResponse(c, 400, "Failed to GET students")
	} else {
		return helpers.SuccessResponse(c, 200, "Success to GET students", students)
	}
}

func GetStudentByID(c *fiber.Ctx) error {
	idStudent, _ := strconv.Atoi(c.Params("id"))
	var studentResp models.StudentDetails

	student, err := services.GetStudentByID(database.DbConnection, idStudent)
	studentResp.ID = student.ID
	studentResp.StudentName = student.StudentName
	studentResp.City = student.City
	studentResp.StartingYear = student.StartingYear

	if err != nil {
		return helpers.ErrorResponse(c, 400, "Failed to GET student")
	}

	courses, err := services.GetCourseTakenDetail(database.DbConnection, idStudent)

	studentResp.CourseTaken = courses

	if err != nil {
		return helpers.ErrorResponse(c, 400, "Failed to GET student")
	} else {
		return helpers.SuccessResponse(c, 200, "Success to GET student", studentResp)
	}
}

func GetCourseTakenByStudent(c *fiber.Ctx) error {
	idStudent, _ := strconv.Atoi(c.Params("id"))

	courses, err := services.GetCourseTakenByStudent(database.DbConnection, idStudent)

	if err != nil {
		return helpers.ErrorResponse(c, 400, "Failed to GET student's courses")
	} else {
		return helpers.SuccessResponse(c, 200, "Success to GET student's courses", courses)
	}
}

func GetCourseTakenDetail(c *fiber.Ctx) error {
	idStudent, _ := strconv.Atoi(c.Params("id"))

	courses, err := services.GetCourseTakenDetail(database.DbConnection, idStudent)

	if err != nil {
		return helpers.ErrorResponse(c, 400, "Failed to GET student's courses")
	} else {
		return helpers.SuccessResponse(c, 200, "Success to GET student's courses", courses)
	}
}
