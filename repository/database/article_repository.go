package database

import (
	"banking_crowd/models/articles"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Save(article articles.Article) (articles.Article, error)
	FindByIdTutor(TutorID int) ([]articles.Article, error)
	FindAll() ([]articles.Article, error)
	Delete(ID int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) Save(article articles.Article) (articles.Article, error) {
	err := r.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *articleRepository) FindByIdTutor(TutorID int) ([]articles.Article, error) {
	var articles []articles.Article
	err := r.db.Where("tutor_id = ? ", TutorID).Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (r *articleRepository) FindAll() ([]articles.Article, error) {
	var articles []articles.Article
	err := r.db.Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (r *articleRepository) Delete(ID int) error {
	article := articles.Article{}
	err := r.db.Where("id = ?", ID).Delete(&article).Error
	if err != nil {
		return err
	}

	return nil
}
