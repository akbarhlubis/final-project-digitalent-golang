package controller

import "github.com/gin-gonic/gin"

func GetArticles(c *gin.Context) {
	// send request to get all articles
	c.JSON(200, gin.H{
		"message": "Get all articles",
	})
}

func GetArticleById(c *gin.Context) {
	// send request to get article by id
	c.JSON(200, gin.H{
		"message": "Get article by id",
	})
}

func UpdateArticleById(c *gin.Context) {
	// send request to update article by id
	c.JSON(200, gin.H{
		"message": "Update article by id",
	})
}

func DeleteArticleById(c *gin.Context) {
	// send request to delete article by id
	c.JSON(200, gin.H{
		"message": "Delete article by id",
	})
}

func CreateArticle(c *gin.Context) {
	// send request to create article
	c.JSON(200, gin.H{
		"message": "Create article",
	})
}

func GetArticleCategories(c *gin.Context) {
	// send request to get article categories
	c.JSON(200, gin.H{
		"message": "Get article categories",
	})
}
