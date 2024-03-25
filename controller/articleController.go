package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/helpers"
	"final-project-akbar/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	db := config.DBInit()
	userData := c.MustGet("user").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	article := model.Article{}
	userID := int(userData["id"].(float64))

	if contentType != "application/json" {
		c.ShouldBindJSON(&article)
	} else {
		c.ShouldBind(&article)

	}

	article.UserID = uint(userID)

	err := db.Debug().Create(&article).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error":   "Failed to create article",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func GetArticleCategories(c *gin.Context) {
	// send request to get article categories
	c.JSON(200, gin.H{
		"message": "Get article categories",
	})
}
