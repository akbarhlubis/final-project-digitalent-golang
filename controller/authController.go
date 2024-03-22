package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	// send request to success login
	c.JSON(200, gin.H{
		"message": "Login success",
	})
}

func Register(c *gin.Context) {
	// send request to success register
	c.JSON(200, gin.H{
		"message": "Register success",
	})
}

func Logout(c *gin.Context) {
	// send request to success logout
	c.JSON(200, gin.H{
		"message": "Logout success",
	})
}
