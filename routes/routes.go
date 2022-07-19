package routes

import (
	createstudent "main/modules/student/useCase/createStudent"
	"main/modules/student/useCase/deleteStudent"
	liststudents "main/modules/student/useCase/listStudents"
	updatestudent "main/modules/student/useCase/updateStudent"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		students := main.Group("students")
		{
			students.GET("/", liststudents.ListStudentsController)
			students.POST("/", createstudent.CreateStudentController)
			students.PUT("/:id", updatestudent.UpdateStudentController)
			students.DELETE("/:id", deleteStudent.DeleteStudentController)
		}
	}

	return router
}
