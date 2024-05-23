package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/helpers"
	"final-project-akbar/model"
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
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unathorized",
			"message": err.Error(),
		})
		return
	}

	User.CheckPassword(password)
	token, err := helpers.GenerateToken(User.ID, User.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unathorized",
			"message": err.Error(),
		})
		return
	}

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

	if err := User.HashPassword(User.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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
	verifyToken, err := helpers.VerifyToken(c)
	_ = verifyToken

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})
}
