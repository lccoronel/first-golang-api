package updatestudent

import (
	"errors"
	"main/database"
	"main/modules/student/entities"
)

func UpdateStudentUseCase(id string, paramStudent entities.Student) (entities.Student, error) {
	var student entities.Student

	amountStudents := len(database.StudentsList)

	for i := 0; i < amountStudents; i++ {
		if database.StudentsList[i].Name != id {
			return student, errors.New("student not found")
		}

		student = database.StudentsList[0]

		student.Name = paramStudent.Name
		student.Age = paramStudent.Age
		student.Gender = paramStudent.Gender

		database.StudentsList = append(database.StudentsList, student)
	}

	return student, nil
}
