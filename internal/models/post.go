package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Topic   string `json:"topic"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (p *Post) TableName() string {
	return "posts"
}
