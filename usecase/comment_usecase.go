package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ICommentUseCase interface {
	GetAllComments() ([]model.CommentResponse, error)
	GetCommentByID(commentId uint) (model.CommentResponse, error)
	GetCommentsByArticleID(articleId uint) ([]model.CommentResponse, error)
	CreateComment(comment model.Comment) (model.CommentResponse, error)
	UpdateComment(comment model.Comment, commentId uint) (model.CommentResponse, error)
	DeleteComment(commentId uint) error
}

type commentUseCase struct {
	cr repository.ICommentRepository
	cv validator.ICommentValidator
}

func NewCommentUseCase(cr repository.ICommentRepository, cv validator.ICommentValidator) ICommentUseCase {
	return &commentUseCase{cr, cv}
}

func (cu *commentUseCase) GetAllComments() ([]model.CommentResponse, error) {
	comments := []model.Comment{}
	commentResponse := []model.CommentResponse{}
	if err := cu.cr.GetAllComments(&comments); err != nil {
		return nil, err
	}
	for _, comment := range comments {
		c := model.CommentResponse{
			ID:        comment.ID,
			ArticleID: comment.ArticleID,
			Author:    comment.Author,
			Reply:     comment.Reply,
			Comment:   comment.Comment,
			CreatedAt: comment.CreatedAt,
		}
		commentResponse = append(commentResponse, c)
	}
	return commentResponse, nil
}

func (cu *commentUseCase) GetCommentByID(commentId uint) (model.CommentResponse, error) {
	comment := model.Comment{}
	if err := cu.cr.GetCommentByID(&comment, commentId); err != nil {
		return model.CommentResponse{}, err
	}
	commentResponse := model.CommentResponse{
		ID:        comment.ID,
		ArticleID: comment.ArticleID,
		Author:    comment.Author,
		Reply:     comment.Reply,
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt,
	}
	return commentResponse, nil
}

func (cu *commentUseCase) GetCommentsByArticleID(articleId uint) ([]model.CommentResponse, error) {
	comments := []model.Comment{}
	commentResponse := []model.CommentResponse{}
	if err := cu.cr.GetCommentsByArticleID(&comments, articleId); err != nil {
		return nil, err
	}
	for _, comment := range comments {
		c := model.CommentResponse{
			ID:        comment.ID,
			ArticleID: comment.ArticleID,
			Author:    comment.Author,
			Reply:     comment.Reply,
			Comment:   comment.Comment,
			CreatedAt: comment.CreatedAt,
		}
		commentResponse = append(commentResponse, c)
	}
	return commentResponse, nil
}

func (cu *commentUseCase) CreateComment(comment model.Comment) (model.CommentResponse, error) {
	// TODO: Validate there is no reply with the same author and comment
	// current: this check only comments but not replies
	commentsForValidate := []model.Comment{}
	if err := cu.cr.GetAllComments(&commentsForValidate); err != nil {
		return model.CommentResponse{}, err
	}
	if err := cu.cv.CommentValidate(comment, &commentsForValidate); err != nil {
		return model.CommentResponse{}, err
	}
	if err := cu.cr.CreateComment(&comment); err != nil {
		return model.CommentResponse{}, err
	}
	commentResponse := model.CommentResponse{
		ID:        comment.ID,
		ArticleID: comment.ArticleID,
		Author:    comment.Author,
		Reply:     comment.Reply,
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt,
	}
	return commentResponse, nil
}

func (cu *commentUseCase) UpdateComment(comment model.Comment, commentId uint) (model.CommentResponse, error) {
	updateTargetComment := model.Comment{}
	if err := cu.cr.GetCommentByID(&updateTargetComment, commentId); err != nil {
		return model.CommentResponse{}, err
	}
	if err := cu.cv.CommentUpdateValidate(comment, updateTargetComment.Author); err != nil {
		return model.CommentResponse{}, err
	}
	if err := cu.cr.UpdateComment(&comment, commentId); err != nil {
		return model.CommentResponse{}, err
	}
	commentResponse := model.CommentResponse{
		ID:        comment.ID,
		ArticleID: comment.ArticleID,
		Author:    comment.Author,
		Reply:     comment.Reply,
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt,
	}
	return commentResponse, nil
}

func (cu *commentUseCase) DeleteComment(commentId uint) error {
	if err := cu.cr.DeleteComment(commentId); err != nil {
		return err
	}
	return nil
}
