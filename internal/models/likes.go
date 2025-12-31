package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	PostID uint `json:"post_id" gorm:"index;uniqueIndex:idx_post_user"`
	UserID uint `json:"user_id" gorm:"index;uniqueIndex:idx_post_user"`
}

func (l *Like) TableName() string {
	return "likes"
}
