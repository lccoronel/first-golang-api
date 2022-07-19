package deleteStudent

import (
	"main/modules/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStudentController(c *gin.Context) {
	id := c.Param("id")

	message, error := DeleteStudentUseCase(id)

	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, dtos.Response{Message: error.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, dtos.Response{Message: message})
	}
}
