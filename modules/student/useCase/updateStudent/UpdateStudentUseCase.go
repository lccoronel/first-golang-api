package updatestudent

import (
	"errors"
	"main/database"
	"main/modules/student/entities"
)

func UpdateStudentUseCase(id string, paramStudent entities.Student) (entities.Student, error) {
	var student entities.Student

	amountStudent := len(database.StudentsList)

	if amountStudent == 0 {
		return student, errors.New("student not found")
	}

	for i := 0; i < amountStudent; i++ {
		if database.StudentsList[i].Id.String() == id {
			student = database.StudentsList[i]

			database.StudentsList[i].Name = paramStudent.Name
			database.StudentsList[i].Age = paramStudent.Age
			database.StudentsList[i].Gender = paramStudent.Gender

		}
	}

	if student.Name == "" {
		return student, errors.New("student not found")
	}

	return student, nil
}
