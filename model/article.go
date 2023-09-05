package model

import "time"

type Article struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Likes     uint16    `gorm:"size:255;not null;default:0" json:"likes"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type ArticleResponse struct {
	ID    uint64 `json:"id" gorm:"primary_key"`
	Title string `json:"title" gorm:"size:255;not null;unique"`
	Likes uint16 `json:"likes" gorm:"size:255;not null;"`
}
