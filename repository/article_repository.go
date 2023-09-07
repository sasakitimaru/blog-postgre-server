package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IArticleRepository interface {
	GetAllArticles(articles *[]model.Article) error
	GetArticleByID(articles *model.Article, articleId uint) error
	CreateArticle(articles *model.Article) error
	UpdateArticle(articles *model.Article, articleId uint) error
	DeleteArticle(articleId uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) GetAllArticles(articles *[]model.Article) error {
	if err := ar.db.Order("created_at").Find(articles).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) GetArticleByID(article *model.Article, articleId uint) error {
	if err := ar.db.Where("id = ?", articleId).First(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) CreateArticle(article *model.Article) error {
	if err := ar.db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) UpdateArticle(article *model.Article, articleId uint) error {
	updateData := map[string]interface{}{
		"title": article.Title,
		"likes": article.Likes,
	}
	if err := ar.db.Model(article).Clauses(clause.Returning{}).Where("id = ?", articleId).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (tr *articleRepository) DeleteArticle(articleId uint) error {
	if err := tr.db.Where("id = ?", articleId).Delete(&model.Article{}).Error; err != nil {
		return err
	}
	return nil
}
