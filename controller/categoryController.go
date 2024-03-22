package controller

import "github.com/gin-gonic/gin"

func GetCategories(c *gin.Context) {
	// send request to get all categories
	c.JSON(200, gin.H{
		"message": "Get all categories",
	})
}

func GetCategoryById(c *gin.Context) {
	// send request to get category by id
	c.JSON(200, gin.H{
		"message": "Get category by id",
	})
}

func UpdateCategoryById(c *gin.Context) {
	// send request to update category by id
	c.JSON(200, gin.H{
		"message": "Update category by id",
	})
}

func DeleteCategoryById(c *gin.Context) {
	// send request to delete category by id
	c.JSON(200, gin.H{
		"message": "Delete category by id",
	})
}

func CreateCategory(c *gin.Context) {
	// send request to create category
	c.JSON(200, gin.H{
		"message": "Create category",
	})
}

func GetCategoryRecipes(c *gin.Context) {
	// send request to get category recipes
	c.JSON(200, gin.H{
		"message": "Get category recipes",
	})
}
