package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id" gorm:"index"`
	Comment string `json:"comment"`
	Author  string `json:"author"`
}
