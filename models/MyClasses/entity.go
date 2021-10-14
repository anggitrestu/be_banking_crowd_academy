package myclasses

import (
	"time"

	"gorm.io/gorm"
)

type MyClass struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	ClassID   int            `json:"class_id"`
	LearnerID int            `json:"learner_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
