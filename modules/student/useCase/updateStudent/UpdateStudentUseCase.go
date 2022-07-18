package updatestudent

import (
	"main/database"
	"main/modules/student/entities"
)

func UpdateStudentUseCase(id string) entities.Student {
	var student entities.Student

	amountStudents := len(database.StudentsList)

	for i := 0; amountStudents >= i; i++ {
		// if database.StudentsList[i].Name == id {
		// 	student = database.StudentsList[i]
		// }
	}

	return student
}
