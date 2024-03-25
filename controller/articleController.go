package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/model"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	var articles []model.Article

	config.DBInit().Find(&articles)

	c.JSON(200, gin.H{
		"message": "Get articles",
		"data":    articles,
	})
}

func GetArticleById(c *gin.Context) {
	var article model.Article

	id := c.Param("id")

	config.DBInit().Where("id = ?", id).Find(&article)

	c.JSON(200, gin.H{
		"message": "Get article by id",
		"data":    article,
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
	var article model.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// send request to create article
	c.JSON(200, gin.H{
		"message": "Create article",
		"data":    article,
	})
}

func GetArticleCategories(c *gin.Context) {
	// send request to get article categories
	c.JSON(200, gin.H{
		"message": "Get article categories",
	})
}
