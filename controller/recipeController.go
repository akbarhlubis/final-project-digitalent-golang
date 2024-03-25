package controller

import (
	"final-project-akbar/config"
	"final-project-akbar/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecipes(c *gin.Context) {
	var recipes []model.Recipe
	config.DBInit().Find(&recipes)
	c.JSON(http.StatusOK, gin.H{
		"recipes": recipes,
	})

}

func GetRecipeById(c *gin.Context) {
	// send request to get recipe by id
	c.JSON(200, gin.H{
		"message": "Get recipe by id",
	})
}

func UpdateRecipeById(c *gin.Context) {
	var recipe model.Recipe

	id := c.Param("id")

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	steps := c.PostForm("steps")

	recipe.Name = name
	recipe.Description = description
	recipe.Steps = steps

	config.DBInit().Where("id = ?", id).Updates(&recipe)

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
	var recipe model.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	name := c.PostForm("name")
	description := c.PostForm("description")
	steps := c.PostForm("steps")

	recipe.Name = name
	recipe.Description = description
	recipe.Steps = steps

	config.DBInit().Create(&recipe)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been created",
		"data":    recipe,
	})
}
