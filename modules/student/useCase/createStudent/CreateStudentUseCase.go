package createstudent

import (
	"errors"
	"main/database"
	"main/modules/student/entities"

	uuid "github.com/satori/go.uuid"
)

func CreateStudentUseCase(student entities.Student) (entities.Student, error) {
	var existsStudent entities.Student

	amountStudents := len(database.StudentsList)

	for i := 0; i < amountStudents; i++ {
		if database.StudentsList[i].Name == student.Name {
			existsStudent = database.StudentsList[i]
		}
	}

	if existsStudent.Name != "" {
		return student, errors.New("student already exists")
	}

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
