package model

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients string   `json:"ingredients"`
	Steps       string   `json:"steps"`
	UserID      uint     `json:"user_id"`
	User        User     `json:"user"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category"`
}
