package createstudent

import (
	"main/database"
	"main/modules/student/entities"
)

func CreateStudentUseCase(student entities.Student) {
	database.StudentsList = append(database.StudentsList, student)
}
