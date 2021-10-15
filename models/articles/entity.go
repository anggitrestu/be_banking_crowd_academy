package articles

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID             int            `gorm:"primaryKey" json:"id"`
	TutorID        int            `json:"tutor_id"`
	Kategori       string            `json:"kategori"`
	Judul          string         `gorm:"size:256" json:"judul"`
	Konten         string         `json:"konten"`
	TanggalPosting string         `json:"tanggal_posting"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
