package database

import (
	myclasses "banking_crowd/models/MyClasses"

	"gorm.io/gorm"
)

type MyClassRepository interface {
	FindAllByLearnerID(learnerID int) ([]myclasses.MyClass, error)
	CheckClass(classID int, learnerID int) (myclasses.MyClass, error)
	Save(myClass myclasses.MyClass) (myclasses.MyClass, error)
}

type myClassRepository struct {
	db *gorm.DB
}

func NewMyClassRepository(db *gorm.DB) *myClassRepository {
	return &myClassRepository{db}
}

func (r *myClassRepository) FindAllByLearnerID(learnerID int) ([]myclasses.MyClass, error) {
	var myClasses []myclasses.MyClass
	err := r.db.Where("learner_id = ?", learnerID).Preload("Class").Find(&myClasses).Error
	if err != nil {
		return myClasses, err
	}

	return myClasses, nil
}

func (r *myClassRepository) CheckClass(classID int, learnerID int) (myclasses.MyClass, error) {
	var myClass myclasses.MyClass
	err := r.db.Where("class_id = ? ", classID).Where("learner_id = ?", learnerID).Find(&myClass).Error
	if err != nil {
		return myClass, err
	}

	return myClass, nil
}
func (r *myClassRepository) Save(myClass myclasses.MyClass) (myclasses.MyClass, error) {
	err := r.db.Create(&myClass).Error
	if err != nil {
		return myClass, err
	}
	return myClass, nil
}
