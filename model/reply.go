package model

import "time"

type Reply struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CommentID uint64    `gorm:"not null" json:"comment_id"`
	Author    string    `gorm:"size:255;not null;default: Anonymous" json:"author"`
	Comment   string    `gorm:"size:255;not null;" json:"comment"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ReplyResponse struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CommentID uint64    `json:"comment_id" gorm:"not null"`
	Author    string    `json:"author" gorm:"size:255;not null;default: Anonymous"`
	Comment   string    `json:"comment" gorm:"size:255;not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
