package deleteStudent

import (
	"errors"
	"main/modules/student/repositories"
)

func DeleteStudentUseCase(id string) (string, error) {
	var message string

	existsStudent := repositories.FindStudentById(id)

	if existsStudent.Name == "" {
		return "", errors.New("student not exists")
	}

	deletedUser := repositories.DeleteStudentById(id)

	if deletedUser {
		message = "deleted user"
	}

	return message, nil
}
