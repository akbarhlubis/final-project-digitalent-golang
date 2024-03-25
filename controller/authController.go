package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/helpers"
	"final-project-akbar/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func Login(c *gin.Context) {
	db := config.DBInit()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}
	password := ""

	if contentType != appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	// get user by email and show in the log
	fmt.Println(User)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unathorized",
			"message": err.Error(),
		})
		return
	}

	User.CheckPassword(password)
	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := model.User{}

	if contentType != appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    User.ID,
		"email": User.Email,
		"name":  User.Name,
	})
}

func Logout(c *gin.Context) {
	// send request to success logout
	c.JSON(200, gin.H{
		"message": "Logout success",
	})
}
