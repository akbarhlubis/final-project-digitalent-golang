package controller

import "github.com/gin-gonic/gin"

func GetUsers(c *gin.Context) {
	// send request to get all users
	c.JSON(200, gin.H{
		"message": "Get all users",
	})
}

func GetUserById(c *gin.Context) {
	// send request to get user by id
	c.JSON(200, gin.H{
		"message": "Get user by id",
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
