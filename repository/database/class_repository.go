package database

import (
	"banking_crowd/models/classes"

	"gorm.io/gorm"
)

type ClassRepository interface {
	Save(class classes.Class) (classes.Class, error)
	FindByIdTutor(TutorID int) ([]classes.ResponseClass, error)
	FindAll() ([]classes.ResponseClass, error)
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

func (r *classRepository) FindByIdTutor(TutorID int) ([]classes.ResponseClass, error) {
	var classes []classes.ResponseClass
	err := r.db.Where("tutor_id = ? ", TutorID).Find(&classes).Error
	if err != nil {
		return classes, err
	}
	return classes, nil
}

func (r *classRepository) FindAll() ([]classes.ResponseClass, error) {
	var classes []classes.ResponseClass
	err := r.db.Raw("select classes.*, learners.email from classes inner join my_classes  on classes.id = my_classes.class_id inner join learners  on my_classes.learner_id = learners.id").Scan(&classes).Error
	if err != nil {
		return classes, err
	}
	return classes, nil
}

// func (r *classRepository) FindAll() ([]classes.Class, error) {
// 	var classes []classes.Class
// 	err := r.db.Preload(clause.Associations).Find(&classes).Error
// 	if err != nil {
// 		return classes, err
// 	}
// 	return classes, nil
// }

/*

select classes.*, learners.email from classes inner join my_classes  on classes.id = my_classes.class_id
inner join learners  on my_classes.learner_id = learners.id;
*/
