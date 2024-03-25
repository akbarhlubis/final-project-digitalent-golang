package model

type Category struct {
	// gorm.Model
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Recipes     []Recipe `json:"recipes"`
}
