package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/db_digitalent_final_project"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB: " + DBName)
	// fmt.Println(dsn)
	return db
}

func GetDB() *gorm.DB {
	return DBInit()
}
