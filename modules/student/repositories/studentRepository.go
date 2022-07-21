package repositories

import (
	"fmt"
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
			helpers.RemoveIndex(i)

			fmt.Println("database => ", database.StudentsList)

			deletedUser = true
		}
	}

	return deletedUser
}

func FindStudentByName(name string) entities.Student {
	var existsStudent entities.Student

	for i := 0; i < len(database.StudentsList); i++ {
		if database.StudentsList[i].Name == name {
			existsStudent = database.StudentsList[i]
		}
	}

	return existsStudent
}
