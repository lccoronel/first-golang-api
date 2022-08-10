package migrations

import (
	"main/modules/student/entities"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entities.Student{})
}
