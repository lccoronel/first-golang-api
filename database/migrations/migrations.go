package migrations

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate()
}
