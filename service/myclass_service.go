package service

import (
	myclasses "banking_crowd/models/MyClasses"
	"banking_crowd/repository/database"
)

type MyClassService interface {
	CreateClass(input myclasses.CreateMyClassInput, learnerID int) (myclasses.MyClass, error)
	GetAllMyClass(learnerID int) ([]myclasses.ResponseMyClass, error)
	IsExistMyClass(input myclasses.CreateMyClassInput, learnerID int) (myclasses.MyClass, error)
}

type myClassService struct {
	repository database.MyClassRepository
}

func NewMyClassService(repository database.MyClassRepository) *myClassService {
	return &myClassService{repository}
}

func (s *myClassService) CreateClass(input myclasses.CreateMyClassInput, learnerID int) (myclasses.MyClass, error) {
	myclass := myclasses.MyClass{}
	myclass.LearnerID = learnerID
	myclass.ClassID = input.ClassID

	newMyCourse, err := s.repository.Save(myclass)
	if err != nil {
		return newMyCourse, err
	}

	return newMyCourse, nil

}

func (s *myClassService) GetAllMyClass(learnerID int) ([]myclasses.ResponseMyClass, error) {
	myclass, err := s.repository.FindAllByLearnerID(learnerID)
	if err != nil {
		return myclass, err
	}
	return myclass, nil
}

func (s *myClassService) IsExistMyClass(input myclasses.CreateMyClassInput, learnerID int) (myclasses.MyClass, error) {
	myclass, err := s.repository.CheckClass(input.ClassID, learnerID)
	if err != nil {
		return myclass, err
	}
	return myclass, nil
}
