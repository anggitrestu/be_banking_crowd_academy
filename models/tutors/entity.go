package tutors

import (
	"banking_crowd/models/articles"
	"time"

	"gorm.io/gorm"
)

type Tutor struct {
	ID         int              `gorm:"primaryKey" json:"id"`
	Nama       string           `gorm:"size:256" json:"nama"`
	MasaKerja  string           `gorm:"size:256" json:"masa_kerja"`
	SitusWeb   string           `gorm:"size:256" json:"situs_web"`
	Kompetensi string           `json:"kompetensi"`
	Pekerjaan  string           `gorm:"size:256" json:"pekerjaan"`
	Password   string           `gorm:"size:256" json:"password"`
	Email      string           `gorm:"unique" json:"email"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	DeletedAt  gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
	Articles   articles.Article `gorm:"foreignKey:TutorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"articles"`
}
