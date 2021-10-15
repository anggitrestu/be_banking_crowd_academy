package database

import (
	"banking_crowd/models/classes"

	"gorm.io/gorm"
)

type ClassRepository interface {
	Save(class classes.Class) (classes.Class, error)
	FindByIdTutor(TutorID int) ([]classes.Class, error)
	FindAll() ([]classes.Class, error)
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *classRepository {
	return &classRepository{db}
}

func (r *classRepository) Save(class classes.Class) (classes.Class, error) {
	err := r.db.Create(&class).Error
	if err != nil {
		return class, err
	}

	return class, nil
}

func (r *classRepository) FindByIdTutor(TutorID int) ([]classes.Class, error) {
	var classes []classes.Class
	err := r.db.Where("tutor_id = ? ", TutorID).Find(&classes).Error
	if err != nil {
		return classes, err
	}
	return classes, nil
}

func (r *classRepository) FindAll() ([]classes.Class, error) {
	var classes []classes.Class
	err := r.db.Raw("SELECT * FROM classes").Scan(&classes).Error
	if err != nil {
		return classes, err
	}
	return classes, nil
}
