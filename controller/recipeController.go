package controller

import "github.com/gin-gonic/gin"

func GetRecipes(c *gin.Context) {
	// send request to get all recipes
	c.JSON(200, gin.H{
		"message": "Get all recipes",
	})
}

func GetRecipeById(c *gin.Context) {
	// send request to get recipe by id
	c.JSON(200, gin.H{
		"message": "Get recipe by id",
	})
}

func UpdateRecipeById(c *gin.Context) {
	// send request to update recipe by id
	c.JSON(200, gin.H{
		"message": "Update recipe by id",
	})
}

func DeleteRecipeById(c *gin.Context) {
	// send request to delete recipe by id
	c.JSON(200, gin.H{
		"message": "Delete recipe by id",
	})
}

func CreateRecipe(c *gin.Context) {
	// send request to create recipe
	c.JSON(200, gin.H{
		"message": "Create recipe",
	})
}
