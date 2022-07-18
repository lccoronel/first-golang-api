package updatestudent

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudentController(c *gin.Context) {
	id := c.Param("id")

	student := UpdateStudentUseCase(id)

	c.IndentedJSON(http.StatusOK, student)
}
