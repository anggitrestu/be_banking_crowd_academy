package learners

import (
	"time"

	"gorm.io/gorm"
)

type Learner struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Nama      string         `gorm:"size:256" json:"nama"`
	Usia      int            `json:"usia"`
	Pekerjaan string         `gorm:"size:256" json:"pekerjaan"`
	Password  string         `gorm:"size:256" json:"password"`
	Email     string         `gorm:"unique" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
