package updatestudent

import (
	"main/modules/dtos"
	"main/modules/student/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudentController(c *gin.Context) {
	id := c.Param("id")
	var student entities.Student

	if err := c.BindJSON(&student); err != nil {
		return
	}

	studentResponse, error := UpdateStudentUseCase(id, student)

	if error == nil {
		c.IndentedJSON(http.StatusOK, studentResponse)
	} else {
		c.IndentedJSON(http.StatusNotFound, dtos.Response{Message: error.Error()})
	}
}
