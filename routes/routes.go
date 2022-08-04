package routes

import (
	createStudent "main/modules/student/useCase/createStudent"
	"main/modules/student/useCase/deleteStudent"
	listStudents "main/modules/student/useCase/listStudents"
	updateStudent "main/modules/student/useCase/updateStudent"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		students := main.Group("students")
		{
			students.GET("/", listStudents.ListStudentsController)
			students.POST("/", createStudent.CreateStudentController)
			students.PUT("/:id", updateStudent.UpdateStudentController)
			students.DELETE("/:id", deleteStudent.DeleteStudentController)
		}
	}

	return router
}
