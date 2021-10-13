package service

import (
	"banking_crowd/models/tutors"
	"banking_crowd/repository/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type TutorService interface {
	RegisterTutor(input tutors.RegisterUserInput) (tutors.Tutor, error)
	GetTutorByID(ID int) (tutors.Tutor, error)
	Login(input tutors.LogisUserInput) (tutors.Tutor, error)
}

type tutorService struct {
	repository database.TutorRepository
}

func NewTutorService(repository database.TutorRepository) *tutorService {
	return &tutorService{repository}
}

func (s *tutorService) RegisterTutor(input tutors.RegisterUserInput) (tutors.Tutor, error) {
	tutor := tutors.Tutor{}
	tutor.Nama = input.Nama
	tutor.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return tutor, err
	}
	tutor.Password = string(password)

	newTutor, err := s.repository.Save(tutor)
	if err != nil {
		return newTutor, err
	}

	return newTutor, nil

}

func (s *tutorService) GetTutorByID(ID int) (tutors.Tutor, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found with that ID")
	}
	return user, nil

}

func (s *tutorService) Login(input tutors.LogisUserInput) (tutors.Tutor, error) {
	email := input.Email
	password := input.Password

	tutor, err := s.repository.FindByEmail(email)
	if err != nil {
		return tutor, nil
	}

	if tutor.ID == 0 {
		return tutor, errors.New("no tutor found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(tutor.Password), []byte(password))
	if err != nil {
		return tutor, err
	}

	return tutor, nil

}
