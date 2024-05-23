package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecipes(c *gin.Context) {
	var recipes []model.Recipe     // create a slice of recipe
	config.DBInit().Find(&recipes) // find all recipes

	if len(recipes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No recipes found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get all recipes",
		"data":    recipes, // return all recipes
	})

}

func GetRecipeById(c *gin.Context) {
	// send request to get recipe by id
	var recipe model.Recipe
	// get the id from the url
	id := c.Param("id")
	// find the recipe by id in the database
	config.DBInit().Where("id = ?", id).Find(&recipe)

	// if the recipe is not found
	if recipe.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Recipe not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get recipe by id",
		"data":    recipe,
	})
}

func UpdateRecipeById(c *gin.Context) {
	var recipe model.Recipe

	id := c.Param("id")

	if err := config.DBInit().Where("id = ?", id).First(&recipe).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Recipe not found",
		})
		return

	}

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	config.DBInit().Save(&recipe)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been updated",
		"data":    recipe,
	})
}

func DeleteRecipeById(c *gin.Context) {
	var recipe model.Recipe

	id := c.Param("id")

	config.DBInit().Where("id = ?", id).Delete(&recipe)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted",
	})
}

func CreateRecipe(c *gin.Context) {
	var recipe model.Recipe // declare a recipe struct

	// condition if the request is not a json
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// create the recipe
	if err := config.DBInit().Create(&recipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been created",
		"data":    recipe,
	})
}
