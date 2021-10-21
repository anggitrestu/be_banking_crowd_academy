package database

import (
	myclasses "banking_crowd/models/MyClasses"

	"gorm.io/gorm"
)

type MyClassRepository interface {
	FindAllByLearnerID(learnerID int) ([]myclasses.ResponseMyClass, error)
	CheckClass(classID int, learnerID int) (myclasses.MyClass, error)
	Save(myClass myclasses.MyClass) (myclasses.MyClass, error)
}

type myClassRepository struct {
	db *gorm.DB
}

func NewMyClassRepository(db *gorm.DB) *myClassRepository {
	return &myClassRepository{db}
}

func (r *myClassRepository) FindAllByLearnerID(learnerID int) ([]myclasses.ResponseMyClass, error) {
	var myClasses []myclasses.ResponseMyClass
	err := r.db.Raw("select my_classes.* , classes.* from my_classes inner join classes on my_classes.class_id = classes.id where my_classes.learner_id = ? ", learnerID).Scan(&myClasses).Error
	if err != nil {
		return myClasses, err
	}

	return myClasses, nil
}

// select my_classes.* , classes.* from my_classes inner join classes on my_classes.class_id = classes.id where my_classes.learner_id = 5;

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
