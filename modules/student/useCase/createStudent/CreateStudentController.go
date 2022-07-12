package createstudent

import (
	"main/modules/student/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudentController(c *gin.Context) {
	var student entities.Student

	if err := c.BindJSON(&student); err != nil {
		return
	}

	CreateStudentUseCase(student)

	c.IndentedJSON(http.StatusCreated, student)
}
