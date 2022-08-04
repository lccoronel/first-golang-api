package listStudents

import (
	"main/database"
	"main/modules/student/entities"
)

func ListStudentsUseCase() []entities.Student {
	return database.StudentsList
}
