package model

import "time"

type Comment struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	ArticleID uint64    `gorm:"not null" json:"article_id"`
	Author    string    `gorm:"size:255;not null;default: Anonymous" json:"author"`
	Comment   string    `gorm:"size:255;not null;" json:"comment"`
	Reply     []Reply   `gorm:"foreignkey:CommentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reply"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CommentResponse struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	ArticleID uint64    `json:"article_id" gorm:"not null"`
	Author    string    `json:"author" gorm:"size:255;not null;default: Anonymous"`
	Reply     []Reply   `gorm:"foreignkey:CommentID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"reply"`
	Comment   string    `json:"comment" gorm:"size:255;not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
