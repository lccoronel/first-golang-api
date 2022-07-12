package liststudents

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListStudentsController(c *gin.Context) {
	listStudents := ListStudentsUseCase()

	c.IndentedJSON(http.StatusOK, listStudents)
}
