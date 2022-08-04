package updateStudent

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

	student = repositories.UpdateStudentById(id, paramStudent)

	if student.Name == "" {
		return student, errors.New("student not found")
	}

	return student, nil
}
