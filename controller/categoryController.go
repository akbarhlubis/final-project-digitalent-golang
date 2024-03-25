package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []model.Category
	config.DBInit().Find(&categories)
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func GetCategoryById(c *gin.Context) {
	var category model.Category

	id := c.Param("id")

	config.DBInit().Where("id = ?", id).Find(&category)

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func UpdateCategoryById(c *gin.Context) {
	// send request to update category by id
	c.JSON(200, gin.H{
		"message": "Update category by id",
	})
}

func DeleteCategoryById(c *gin.Context) {
	var category model.Category

	id := c.Param("id")

	config.DBInit().Where("id = ?", id).Delete(&category)

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been deleted",
	})
}

func CreateCategory(c *gin.Context) {
	var category model.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DBInit().Create(&category)

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been created",
		"data":    category,
	})
}

func GetCategoryRecipes(c *gin.Context) {
	// send request to get category recipes
	c.JSON(200, gin.H{
		"message": "Get category recipes",
	})
}
