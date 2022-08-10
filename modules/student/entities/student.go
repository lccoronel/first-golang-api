package entities

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	Gender    bool           `json:"gender"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}
