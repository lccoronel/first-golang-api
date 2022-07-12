package routes

import (
	createstudent "main/modules/student/useCase/createStudent"
	liststudents "main/modules/student/useCase/listStudents"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		students := main.Group("students")
		{
			students.GET("/", liststudents.ListStudentsController)
			students.POST("/", createstudent.CreateStudentController)
		}
	}

	return router
}
