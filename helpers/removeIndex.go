package helpers

import (
	"main/database"
	"main/modules/student/entities"
)

func RemoveIndex(index int) {
	copy(database.StudentsList[index:], database.StudentsList[index+1:])
	database.StudentsList[len(database.StudentsList)-1] = entities.Student{}
	database.StudentsList = database.StudentsList[:len(database.StudentsList)-1]
}
