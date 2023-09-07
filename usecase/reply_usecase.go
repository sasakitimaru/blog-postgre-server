package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IReplyUseCase interface {
	GetAllReplies() ([]model.ReplyResponse, error)
	GetReplyByID(replyId uint) (model.ReplyResponse, error)
	GetRepliesByCommentID(commentId uint) ([]model.ReplyResponse, error)
	CreateReply(reply model.Reply) (model.ReplyResponse, error)
	UpdateReply(reply model.Reply, replyId uint) (model.ReplyResponse, error)
	DeleteReply(replyId uint) error
}

type replyUseCase struct {
	rr repository.IReplyRepository
	cv validator.ICommentValidator
}

func NewReplyUseCase(rr repository.IReplyRepository, cv validator.ICommentValidator) IReplyUseCase {
	return &replyUseCase{rr, cv}
}

func (ru *replyUseCase) GetAllReplies() ([]model.ReplyResponse, error) {
	replies := []model.Reply{}
	replyResponse := []model.ReplyResponse{}
	if err := ru.rr.GetAllReplies(&replies); err != nil {
		return nil, err
	}
	for _, reply := range replies {
		r := model.ReplyResponse{
			ID:        reply.CommentID,
			CommentID: reply.CommentID,
			Author:    reply.Author,
			Comment:   reply.Comment,
			CreatedAt: reply.CreatedAt,
		}
		replyResponse = append(replyResponse, r)
	}
	return replyResponse, nil
}

func (ru *replyUseCase) GetReplyByID(replyId uint) (model.ReplyResponse, error) {
	reply := model.Reply{}
	if err := ru.rr.GetReplyByID(&reply, replyId); err != nil {
		return model.ReplyResponse{}, err
	}
	replyResponse := model.ReplyResponse{
		ID:        reply.CommentID,
		CommentID: reply.CommentID,
		Author:    reply.Author,
		Comment:   reply.Comment,
		CreatedAt: reply.CreatedAt,
	}
	return replyResponse, nil
}

func (ru *replyUseCase) GetRepliesByCommentID(commentId uint) ([]model.ReplyResponse, error) {
	replies := []model.Reply{}
	replyResponse := []model.ReplyResponse{}
	if err := ru.rr.GetRepliesByCommentID(&replies, commentId); err != nil {
		return nil, err
	}
	for _, reply := range replies {
		r := model.ReplyResponse{
			ID:        reply.CommentID,
			CommentID: reply.CommentID,
			Author:    reply.Author,
			Comment:   reply.Comment,
			CreatedAt: reply.CreatedAt,
		}
		replyResponse = append(replyResponse, r)
	}
	return replyResponse, nil
}

func (ru *replyUseCase) CreateReply(reply model.Reply) (model.ReplyResponse, error) {
	// TODO: Validate there is no reply with the same author and comment
	if err := ru.rr.CreateReply(&reply); err != nil {
		return model.ReplyResponse{}, err
	}
	replyResponse := model.ReplyResponse{
		ID:        reply.CommentID,
		CommentID: reply.CommentID,
		Author:    reply.Author,
		Comment:   reply.Comment,
		CreatedAt: reply.CreatedAt,
	}
	return replyResponse, nil
}

func (ru *replyUseCase) UpdateReply(reply model.Reply, replyId uint) (model.ReplyResponse, error) {
	if err := ru.rr.UpdateReply(&reply, replyId); err != nil {
		return model.ReplyResponse{}, err
	}
	replyResponse := model.ReplyResponse{
		ID:        reply.CommentID,
		CommentID: reply.CommentID,
		Author:    reply.Author,
		Comment:   reply.Comment,
		CreatedAt: reply.CreatedAt,
	}
	return replyResponse, nil
}

func (ru *replyUseCase) DeleteReply(replyId uint) error {
	if err := ru.rr.DeleteReply(replyId); err != nil {
		return err
	}
	return nil
}
