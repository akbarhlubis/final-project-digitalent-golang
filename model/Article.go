package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	Bookmarks []Bookmark
}
