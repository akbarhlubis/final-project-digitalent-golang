package model

type Article struct {
	// gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	Bookmarks []Bookmark
}
