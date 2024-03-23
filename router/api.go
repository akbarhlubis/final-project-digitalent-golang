package router

import (
	"final-project-akbar/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Testing route - Single route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Wasap gaesss, pake GET Request!",
		})
	})
	router.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Wasap gaesss, pake POST Request!",
		})
	})

	// API route - Grouping route
	api := router.Group("/api")
	{
		// Auth route
		api.POST("/login", controller.Login)
		api.POST("/register", controller.Register)
		api.POST("/logout", controller.Logout)

		// User route
		api.GET("/users", controller.GetUsers)
		user := api.Group("/user")
		{
			user.GET("/:id", controller.GetUserById)
			user.DELETE("/:id", controller.DeleteUserById)
			user.PUT("/:id", controller.UpdateUserById)
		}

		// Recipe route
		api.GET("/recipes", controller.GetRecipes)
		recipe := api.Group("/recipe")
		{
			recipe.GET("/:id", controller.GetRecipeById)
			recipe.DELETE("/:id", controller.DeleteRecipeById)
			recipe.PUT("/:id", controller.UpdateRecipeById)
			recipe.POST("/", controller.CreateRecipe)
		}

		// Category route
		api.GET("/categories", controller.GetCategories)
		category := api.Group("/category")
		{
			category.GET("/:id", controller.GetCategoryById)
			category.DELETE("/:id", controller.DeleteCategoryById)
			category.PUT("/:id", controller.UpdateCategoryById)
			category.POST("/", controller.CreateCategory)
			category.GET("/:id/recipes", controller.GetCategoryRecipes)
		}

		// Article route
		api.GET("/articles", controller.GetArticles)
		article := api.Group("/article")
		{
			article.GET("/:id", controller.GetArticleById)
			article.DELETE("/:id", controller.DeleteArticleById)
			article.PUT("/:id", controller.UpdateArticleById)
			article.POST("/", controller.CreateArticle)
			article.GET("/categories", controller.GetArticleCategories)
		}
	}

	return router
}
