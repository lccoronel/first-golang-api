package updatestudent

import (
	"errors"
	"main/database"
	"main/modules/student/entities"
	"main/modules/student/repositories"
)

func UpdateStudentUseCase(id string, paramStudent entities.Student) (entities.Student, error) {
	var student entities.Student
	var existsStudent entities.Student

	amountStudent := len(database.StudentsList)

	if amountStudent == 0 {
		return student, errors.New("student not found")
	}

	existsStudent = repositories.FindStudentByName(paramStudent.Name)

	if existsStudent.Name != "" {
		return student, errors.New("student already exists")
	}

	// transfer this action to repositories
	for i := 0; i < amountStudent; i++ {
		if database.StudentsList[i].Id.String() == id {
			database.StudentsList[i].Name = paramStudent.Name
			database.StudentsList[i].Age = paramStudent.Age
			database.StudentsList[i].Gender = paramStudent.Gender

			student = database.StudentsList[i]
		}
	}

	if student.Name == "" {
		return student, errors.New("student not found")
	}

	return student, nil
}
