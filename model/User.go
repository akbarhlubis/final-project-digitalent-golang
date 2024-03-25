package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

type User struct {
	// gorm.Model
	ID        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" form:"name" validate:"required"`
	Username  *string    `json:"username" form:"username"`                                            // *string berarti nullable
	Password  string     `json:"password" form:"password" validate:"required,min=6"`                  // form:"password" berarti ketika melakukan binding dari form, maka field ini akan di bind dengan field password
	Email     string     `json:"email" form:"email" validate:"required,email"`                        // validate:"required,email" berarti field ini wajib diisi dan harus berupa email
	Role      string     `json:"role" form:"role" gorm:"type:ENUM('admin', 'member');default:member"` // gorm:"default:member" berarti ketika membuat user, maka role akan di set default menjadi member
	Recipes   []Recipe   `json:"recipes"`
	Articles  []Article  `json:"articles"`
	Bookmarks []Bookmark `json:"bookmarks"`
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
