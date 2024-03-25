package model

type Bookmark struct {
	// gorm.Model
	ID        uint    `json:"id" gorm:"primaryKey"`
	UserID    uint    `json:"user_id"`
	ArticleID uint    `json:"article_id"`
	RecipeID  uint    `json:"recipe_id"`
	Article   Article `json:"article"`
	Recipe    Recipe  `json:"recipe"`
	User      User    `json:"user"`
}
