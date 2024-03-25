package config

import (
	"final-project-akbar/model"
	"fmt"
)

func SyncDB() {
	db := DBInit()
	db.Debug().AutoMigrate(&model.User{}, &model.Recipe{}, &model.Category{}, &model.Bookmark{}, &model.Article{})
	// db.Debug().AutoMigrate(&model.User{})
	fmt.Println("Database Migrated...")
}
