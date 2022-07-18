package createstudent

import (
	"errors"
	"main/database"
	"main/modules/student/entities"
)

func CreateStudentUseCase(student entities.Student) (entities.Student, error) {
	amountStudents := len(database.StudentsList)

	if amountStudents > 0 {
		for i := 0; i < amountStudents; i++ {
			if database.StudentsList[i].Name == student.Name {
				return student, errors.New("student already exists")
			}
		}
	}

	database.StudentsList = append(database.StudentsList, student)

	return student, nil
}
