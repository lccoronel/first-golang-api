package repositories

import (
	"main/database"
	"main/helpers"
	"main/modules/student/entities"
)

func FindStudentById(id string) entities.Student {
	var existsStudent entities.Student

	for i := 0; i < len(database.StudentsList); i++ {
		if database.StudentsList[i].Id.String() == id {
			existsStudent = database.StudentsList[i]
		}
	}

	return existsStudent
}

func DeleteStudentById(id string) bool {
	var deletedUser bool = false

	// keep developing delete user
	for i := 0; i < len(database.StudentsList); i++ {
		if database.StudentsList[i].Id.String() == id {
			helpers.RemoveIndex(database.StudentsList, i)
			deletedUser = true
		}
	}

	return deletedUser
}
