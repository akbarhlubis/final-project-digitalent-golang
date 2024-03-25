package controller

import "github.com/gin-gonic/gin"

func GetBookmarkById(c *gin.Context) {
	// send request to get bookmark by id
	c.JSON(200, gin.H{
		"message": "Get bookmark by id",
	})
}

func UpdateBookmarkById(c *gin.Context) {
	// send request to update bookmark by id
	c.JSON(200, gin.H{
		"message": "Update bookmark by id",
	})
}

func DeleteBookmarkById(c *gin.Context) {
	// send request to delete bookmark by id
	c.JSON(200, gin.H{
		"message": "Delete bookmark by id",
	})
}

func CreateBookmark(c *gin.Context) {
	// send request to create bookmark
	c.JSON(200, gin.H{
		"message": "Create bookmark",
	})
}
