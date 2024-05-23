package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	config.DBInit().Find(&users) // find all users

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No users found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get all users",
		"data":    users, // return all recipes
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
	var user model.User
	id := c.Param("id")

	if result := config.DBInit().First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := config.DBInit().Save(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if result := config.DBInit().First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := config.DBInit().Delete(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
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

	// username := c.PostForm("username")
	// name := c.PostForm("name")
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	// role := c.PostForm("role")

	// user.Username = new(string)
	// *user.Username = username
	// user.Name = name
	// user.Email = email
	// user.Password = password
	// user.Role = role

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
