package updatestudent

import (
	"main/modules/student/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudentController(c *gin.Context) {
	id := c.Param("id")
	var paramStudent entities.Student

	student, error := UpdateStudentUseCase(id, paramStudent)

	c.IndentedJSON(http.StatusOK, student)
	c.IndentedJSON(http.StatusNotFound, error)
}
