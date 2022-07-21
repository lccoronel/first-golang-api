package helpers

import (
	"fmt"
	"main/database"
	"main/modules/student/entities"
)

func RemoveIndex(index int) []entities.Student {
	fmt.Println("remove", database.StudentsList[:index])
	fmt.Println("rest", database.StudentsList[index+1:])
	return append(database.StudentsList[:index], database.StudentsList[index+1:]...)
}
