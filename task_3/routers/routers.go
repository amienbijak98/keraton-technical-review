package routers

import (
	"keraton-task-3/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello! This is golang mini api project created by Bijak Amien Muttaqie. For submitting Keraton Tech Review Task No. 3")
	})

	//Grouping Routes
	api := app.Group("/task3/api/v1")

	//Student Route
	api.Get("/students", controllers.GetAllStudent)                          //all student
	api.Get("/student/:id", controllers.GetStudentByID)                      //student by id
	api.Get("/student/:id/courses", controllers.GetCourseTakenByStudent)     //courses taken by student id
	api.Get("/student/:id/courses-detail", controllers.GetCourseTakenDetail) //courses taken by student id (detailed)

}
