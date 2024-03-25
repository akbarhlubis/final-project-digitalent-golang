package model

type Recipe struct {
	// gorm.Model
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients string   `json:"ingredients"`
	Steps       string   `json:"steps"`
	UserID      uint     `json:"user_id"`
	User        User     `json:"user"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category"`
	Bookmarks   []Bookmark
}
