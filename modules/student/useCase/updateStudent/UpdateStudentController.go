package updatestudent

import (
	"main/modules/dtos"
	"main/modules/student/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudentController(c *gin.Context) {
	id := c.Param("id")
	var paramStudent entities.Student

	student, error := UpdateStudentUseCase(id, paramStudent)

	if error == nil {
		c.IndentedJSON(http.StatusOK, student)
	} else {
		c.IndentedJSON(http.StatusNotFound, dtos.Error{Message: error.Error()})
	}
}
