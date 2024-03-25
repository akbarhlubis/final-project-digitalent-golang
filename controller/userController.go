package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	if result := config.DBInit().Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if result := config.DBInit().First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUserById(c *gin.Context) {
	// send request to update user by id
	c.JSON(200, gin.H{
		"message": "Update user by id",
	})
}

func DeleteUserById(c *gin.Context) {
	// send request to delete user by id
	c.JSON(200, gin.H{
		"message": "Delete user by id",
	})
}

func CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	username := c.PostForm("username")
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role := c.PostForm("role")

	user.Username = new(string)
	*user.Username = username
	user.Name = name
	user.Email = email
	user.Password = password
	user.Role = role

	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := config.DBInit().Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
