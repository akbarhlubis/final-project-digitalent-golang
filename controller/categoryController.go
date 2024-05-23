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

	if len(categories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No categories found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func GetCategoryById(c *gin.Context) {
	var category model.Category

	id := c.Param("id")

	config.DBInit().Where("id = ?", id).Find(&category)

	if category.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func UpdateCategoryById(c *gin.Context) {
	var category model.Category

	id := c.Param("id")

	if err := config.DBInit().Where("id = ?", id).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DBInit().Save(&category)

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been updated",
		"data":    category,
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

	// create form data
	// name := c.PostForm("name")
	// description := c.PostForm("description")

	// category.Name = name
	// category.Description = description

	// if err := config.DBInit().Create(&category).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

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
