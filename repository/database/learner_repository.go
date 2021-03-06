package database

import (
	"banking_crowd/models/learners"

	"gorm.io/gorm"
)

type LearnerRepository interface {
	Save(learner learners.Learner) (learners.Learner, error)
	FindByID(ID int) (learners.Learner, error)
	Update(learner learners.Learner) (learners.Learner, error)
	FindByEmail(email string) (learners.Learner, error)
	GetLearnerByIdCLass(classID int) ([]learners.Learner, error)
}

type learnerRepository struct {
	db *gorm.DB
}

func NewLearnerRepository(db *gorm.DB) *learnerRepository {
	return &learnerRepository{db}
}

func (r *learnerRepository) Save(learner learners.Learner) (learners.Learner, error) {
	err := r.db.Create(&learner).Error
	if err != nil {
		return learner, err
	}

	return learner, nil
}

func (r *learnerRepository) FindByID(ID int) (learners.Learner, error) {
	var learner learners.Learner
	err := r.db.Where("id = ?", ID).Find(&learner).Error
	if err != nil {
		return learner, err
	}
	return learner, nil
}

func (r *learnerRepository) Update(learner learners.Learner) (learners.Learner, error) {
	err := r.db.Save(&learner).Error
	if err != nil {
		return learner, err
	}

	return learner, nil
}

func (r *learnerRepository) FindByEmail(email string) (learners.Learner, error) {
	var learner learners.Learner
	err := r.db.Where("email = ?", email).Find(&learner).Error
	if err != nil {
		return learner, err
	}

	return learner, nil
}

func (r learnerRepository) GetLearnerByIdCLass(classID int) ([]learners.Learner, error) {
	var learners []learners.Learner
	err := r.db.Raw("select learners.email from learners inner join my_classes on learners.id = my_classes.learner_id where my_classes.class_id = ?;", classID).Scan(&learners).Error
	if err != nil {
		return learners, err
	}
	return learners, nil

}

/*
select learners.email from learners
inner join my_classes on learners.id =
my_classes.learner_id
where my_classes.class_id = ?;
*/
