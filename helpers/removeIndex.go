package helpers

import "main/modules/student/entities"

func RemoveIndex(studentList []entities.Student, index int) []entities.Student {
	return append(studentList[:index], studentList[index+1:]...)
}
