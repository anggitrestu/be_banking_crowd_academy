package classes

import (
	"banking_crowd/models/tutors"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	TutorID   int            `json:"tutor_id"`
	Topik     string         `json:"topik"`
	Jenis     string         `json:"jenis" gorm:"type:enum('gratis', 'berbayar')"`
	Judul     string         `json:"judul" gorm:"size:256"`
	Jadwal    string         `json:"jadwal" gorm:"size:256"`
	LinkZoom  string         `json:"link_zoom" gorm:"size:256"`
	Deskripsi string         `json:"deskripsi"`
	Modul     string         `json:"modul"  gorm:"size:256"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Tutor     tutors.Tutor   `gorm:"foreignKey:TutorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"tutor"`
}
