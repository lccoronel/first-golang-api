package createStudent

import (
	"errors"
	"main/modules/student/entities"
	"main/modules/student/repositories"
)

func CreateStudentUseCase(student entities.Student) (entities.Student, error) {
	var existsStudent entities.Student

	existsStudent = repositories.FindStudentByName(student.Name)

	if existsStudent.Name != "" {
		return student, errors.New("student already exists")
	}

	repositories.CreateStudent(student)

	return student, nil
}
