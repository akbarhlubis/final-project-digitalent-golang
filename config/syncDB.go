package config

import (
	"final-project-akbar/model"
	"fmt"
)

func SyncDB() {
	db := DBInit()
	db.AutoMigrate(&model.User{}, &model.Recipe{}, &model.Category{}, &model.Bookmark{}, &model.Article{})
	fmt.Println("Database Migrated...")
}
