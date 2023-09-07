package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IReplyRepository interface {
	GetAllReplies(replies *[]model.Reply) error
	GetReplyByID(reply *model.Reply, replyId uint) error
	GetRepliesByCommentID(replies *[]model.Reply, commentId uint) error
	CreateReply(reply *model.Reply) error
	UpdateReply(reply *model.Reply, replyId uint) error
	DeleteReply(replyId uint) error
}

type replyRepository struct {
	db *gorm.DB
}

func NewReplyRepository(db *gorm.DB) IReplyRepository {
	return &replyRepository{db}
}

func (rr *replyRepository) GetAllReplies(replies *[]model.Reply) error {
	if err := rr.db.Order("created_at").Find(replies).Error; err != nil {
		return err
	}
	return nil
}

func (rr *replyRepository) GetReplyByID(reply *model.Reply, replyId uint) error {
	if err := rr.db.Where("id = ?", replyId).First(reply).Error; err != nil {
		return err
	}
	return nil
}

func (rr *replyRepository) GetRepliesByCommentID(replies *[]model.Reply, commentId uint) error {
	if err := rr.db.Where("comment_id = ?", commentId).Find(replies).Error; err != nil {
		return err
	}
	return nil
}

func (rr *replyRepository) CreateReply(reply *model.Reply) error {
	if err := rr.db.Create(reply).Error; err != nil {
		return err
	}
	return nil
}

func (rr *replyRepository) UpdateReply(reply *model.Reply, replyId uint) error {
	updateData := map[string]interface{}{
		"comment": reply.Comment,
	}
	if err := rr.db.Model(reply).Where("id = ?", replyId).Updates(updateData).Error; err != nil {
		return err
	}
	return nil
}

func (rr *replyRepository) DeleteReply(replyId uint) error {
	if err := rr.db.Where("id = ?", replyId).Delete(&model.Reply{}).Error; err != nil {
		return err
	}
	return nil
}
