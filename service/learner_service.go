package service

import (
	"banking_crowd/models/learners"
	"banking_crowd/models/tutors"
	"banking_crowd/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LearnerService interface {
	RegisterLearner(input tutors.RegisterUserInput) (learners.Learner, error)
	GetLearnerByID(ID int) (learners.Learner, error)
	LoginLearner(input tutors.LogisUserInput) (learners.Learner, error)
	UpdateLearner(inputID learners.GetLearnerInput, inputData learners.CreateLearnerInput) (learners.Learner, error)
	GetLearnerByIdCLass(classID int) ([]learners.Learner, error)
}

type learnerService struct {
	repository database.LearnerRepository
}

func NewLeranerService(repository database.LearnerRepository) *learnerService {
	return &learnerService{repository}
}

func (s *learnerService) RegisterLearner(input tutors.RegisterUserInput) (learners.Learner, error) {
	learner := learners.Learner{}
	learner.Nama = input.Nama
	learner.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return learner, err
	}
	learner.Password = string(password)

	newLearner, err := s.repository.Save(learner)
	if err != nil {
		return newLearner, err
	}

	return newLearner, nil

}

func (s *learnerService) GetLearnerByID(ID int) (learners.Learner, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}
	return user, nil

}

func (s *learnerService) LoginLearner(input tutors.LogisUserInput) (learners.Learner, error) {
	email := input.Email
	password := input.Password

	learner, err := s.repository.FindByEmail(email)
	if err != nil {
		return learner, nil
	}

	if learner.ID == 0 {
		return learner, errors.New("no learner found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(learner.Password), []byte(password))
	if err != nil {
		return learner, err
	}

	return learner, nil

}

func (s *learnerService) UpdateLearner(inputID learners.GetLearnerInput, inputData learners.CreateLearnerInput) (learners.Learner, error) {
	learner, err := s.repository.FindByID(inputID.ID)

	if err != nil {
		return learner, err
	}

	learner.Nama = inputData.Nama
	learner.Usia = inputData.Usia
	learner.Pekerjaan = inputData.Pekerjaan
	learner.TopikDiminati = inputData.TopikDiminati

	updateLearner, err := s.repository.Update(learner)
	if err != nil {
		return updateLearner, err
	}

	return updateLearner, nil

}

func (s *learnerService) GetLearnerByIdCLass(classID int) ([]learners.Learner, error) {
	learners, err := s.repository.GetLearnerByIdCLass(classID)
	if err != nil {
		return learners, err
	}

	return learners, nil
}
