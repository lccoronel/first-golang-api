package createstudent

import (
	"errors"
	"main/database"
	"main/modules/student/entities"
	"main/modules/student/repositories"

	uuid "github.com/satori/go.uuid"
)

func CreateStudentUseCase(student entities.Student) (entities.Student, error) {
	var existsStudent entities.Student

	existsStudent = repositories.FindStudentByName(student.Name)

	if existsStudent.Name != "" {
		return student, errors.New("student already exists")
	}

	// transfer this action to repositories
	uuid := uuid.NewV4()
	database.StudentsList = append(
		database.StudentsList,
		entities.Student{
			Id:     uuid,
			Name:   student.Name,
			Age:    student.Age,
			Gender: student.Gender,
		},
	)

	return student, nil
}
