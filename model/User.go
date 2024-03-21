package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Username *string `json:"username"` // *string berarti nullable
	Password string  `json:"password"`
	Email    string  `json:"email"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 adalah cost yang direkomendasikan
	if err != nil {
		return err
	}
	user.Password = string(bytes) // jika tidak terjadi error maka password akan di hash
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) // membandingkan password yang diinputkan dengan password yang sudah di hash
	if err != nil {
		return err
	}
	return nil
}
