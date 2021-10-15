package service

import (
	"banking_crowd/models/articles"
	"banking_crowd/repository/database"
	"errors"
)

type ArticleService interface {
	CreateArticle(input articles.CreateArticleInput) (articles.Article, error)
	GetAll(TutorID int) ([]articles.Article, error)
	Delete(ID int) error
}

type articleService struct {
	repository   database.ArticleRepository
	serviceTutor tutorService
}

func NewArticleService(repository database.ArticleRepository, serviceTutor tutorService) *articleService {
	return &articleService{repository, serviceTutor}
}

func (s *articleService) CreateArticle(input articles.CreateArticleInput) (articles.Article, error) {

	tutor, err := s.serviceTutor.GetTutorByID(input.TutorID)
	if err != nil {
		return articles.Article{}, errors.New("tutor not found")
	}
	if tutor.ID == 0 {
		return articles.Article{}, err
	}

	article := articles.Article{}
	article.TutorID = input.TutorID
	article.Kategori = input.Kategori
	article.Judul = input.Judul
	article.Konten = input.Konten

	newArticle, err := s.repository.Save(article)
	if err != nil {
		return newArticle, err
	}

	return newArticle, nil

}

func (s *articleService) GetAll(TutorID int) ([]articles.Article, error) {

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

func (s *articleService) Delete(ID int) error {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
