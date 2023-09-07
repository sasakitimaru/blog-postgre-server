package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	GetAllComments(comments *[]model.Comment) error
	GetCommentByID(comment *model.Comment, commentId uint) error
	GetCommentsByArticleID(comments *[]model.Comment, articleId uint) error
	CreateComment(comment *model.Comment) error
	UpdateComment(comment *model.Comment, commentId uint) error
	DeleteComment(commentId uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &commentRepository{db}
}

func (cr *commentRepository) GetAllComments(comments *[]model.Comment) error {
	if err := cr.db.Order("created_at").Find(comments).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) GetCommentByID(comment *model.Comment, commentId uint) error {
	if err := cr.db.Where("id = ?", commentId).First(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) GetCommentsByArticleID(comments *[]model.Comment, articleId uint) error {
	if err := cr.db.Where("article_id = ?", articleId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) CreateComment(comment *model.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) UpdateComment(comment *model.Comment, commentId uint) error {
	updateData := map[string]interface{}{
		"comment": comment.Comment,
	}
	if err := cr.db.Model(comment).Where("id = ?", commentId).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) DeleteComment(commentId uint) error {
	if err := cr.db.Where("id = ?", commentId).Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}
