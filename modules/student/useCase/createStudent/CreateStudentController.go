package createStudent

import (
	"main/modules/dtos"
	"main/modules/student/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudentController(c *gin.Context) {
	var student entities.Student

	if err := c.BindJSON(&student); err != nil {
		return
	}

	student, error := CreateStudentUseCase(student)

	if error == nil {
		c.IndentedJSON(http.StatusCreated, student)
	} else {
		c.IndentedJSON(http.StatusBadRequest, dtos.Response{Message: error.Error()})
	}

}
