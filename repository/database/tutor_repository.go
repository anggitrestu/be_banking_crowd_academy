package database

import (
	"banking_crowd/models/tutors"

	"gorm.io/gorm"
)

type TutorRepository interface {
	Save(tutor tutors.Tutor) (tutors.Tutor, error)
	FindByID(ID int) (tutors.Tutor, error)
	Update(tutor tutors.Tutor) (tutors.Tutor, error)
	FindByEmail(email string) (tutors.Tutor, error)
}

type tutorRepository struct {
	db *gorm.DB
}

func NewTutorRepository(db *gorm.DB) *tutorRepository {
	return &tutorRepository{db}
}

func (r *tutorRepository) Save(tutor tutors.Tutor) (tutors.Tutor, error) {
	err := r.db.Create(&tutor).Error
	if err != nil {
		return tutor, err
	}

	return tutor, nil
}

func (r *tutorRepository) FindByID(ID int) (tutors.Tutor, error) {
	var tutor tutors.Tutor
	err := r.db.Where("id = ?", ID).Find(&tutor).Error
	if err != nil {
		return tutor, err
	}
	return tutor, nil
}

func (r *tutorRepository) Update(tutor tutors.Tutor) (tutors.Tutor, error) {
	err := r.db.Save(&tutor).Error
	if err != nil {
		return tutor, err
	}

	return tutor, nil
}

func (r *tutorRepository) FindByEmail(email string) (tutors.Tutor, error) {
	var tutor tutors.Tutor
	err := r.db.Where("email = ?", email).Find(&tutor).Error
	if err != nil {
		return tutor, err
	}

	return tutor, nil
}
