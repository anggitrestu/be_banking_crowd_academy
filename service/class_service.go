package service

import (
	"banking_crowd/models/classes"
	"banking_crowd/repository/database"
	"errors"
)

type ClassService interface {
	CreateClass(input classes.CreateClassInput) (classes.Class, error)
	GetAll(TutorID int) ([]classes.Class, error)
}

type classService struct {
	repository   database.ClassRepository
	serviceTutor tutorService
}

func NewClassService(repository database.ClassRepository, serviceTutor tutorService) *classService {
	return &classService{repository, serviceTutor}
}

func (s *classService) CreateClass(input classes.CreateClassInput) (classes.Class, error) {

	tutor, err := s.serviceTutor.GetTutorByID(input.TutorID)
	if err != nil {
		return classes.Class{}, errors.New("tutor not found")
	}
	if tutor.ID == 0 {
		return classes.Class{}, err
	}

	class := classes.Class{}
	class.TutorID = input.TutorID
	class.Jenis = input.Jenis
	class.Judul = input.Judul
	class.Topik = input.Topik
	class.Jadwal = input.Jadwal
	class.Deskripsi = input.Deskripsi
	class.LinkZoom = input.LinkZoom
	class.Modul = input.Modul

	newClass, err := s.repository.Save(class)
	if err != nil {
		return newClass, err
	}

	return newClass, nil

}

func (s *classService) GetAll(TutorID int) ([]classes.Class, error) {

	if TutorID != 0 {
		classes, err := s.repository.FindByIdTutor(TutorID)
		if err != nil {
			return classes, err
		}

		return classes, nil
	}

	classes, err := s.repository.FindAll()
	if err != nil {
		return classes, err
	}

	return classes, nil

}
