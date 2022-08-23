package repositories

import (
	"main/database"
	"main/modules/student/entities"
)

func FindStudentById(id string) entities.Student {
	var existsStudent entities.Student

	for i := 0; i < len(database.StudentsList); i++ {
		// if database.StudentsList[i].ID.String() == id {
		// 	existsStudent = database.StudentsList[i]
		// }
	}

	return existsStudent
}

func DeleteStudentById(id string) bool {
	var deletedUser bool = false

	for i := 0; i < len(database.StudentsList); i++ {
		// if database.StudentsList[i].ID.String() == id {
		// 	helpers.RemoveIndex(i)

		// 	deletedUser = true
		// }
	}

	return deletedUser
}

func FindStudentByName(name string) entities.Student {
	db := database.GetDatabase()

	var student entities.Student

	err := db.First(&student, name).Error

	if err != nil {
		return student
	}

	return student
}

func UpdateStudentById(id string, paramStudent entities.Student) entities.Student {
	var student entities.Student

	for i := 0; i < len(database.StudentsList); i++ {
		// if database.StudentsList[i].ID.String() == id {
		// 	database.StudentsList[i].Name = paramStudent.Name
		// 	database.StudentsList[i].Age = paramStudent.Age
		// 	database.StudentsList[i].Gender = paramStudent.Gender

		// 	student = database.StudentsList[i]
		// }
	}

	return student
}

func CreateStudent(student entities.Student) {
	// uuid := uuid.NewV4()
	// database.StudentsList = append(
	// 	database.StudentsList,
	// 	entities.Student{
	// 		ID:     uuid,
	// 		Name:   student.Name,
	// 		Age:    student.Age,
	// 		Gender: student.Gender,
	// 	},
	// )
}
