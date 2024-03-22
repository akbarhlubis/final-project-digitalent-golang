package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Recipes     []Recipe `json:"recipes"`
}