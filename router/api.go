package router

import (
	"final-project-akbar/controller"
	"final-project-akbar/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Testing route - Single route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Berhasil, pake GET Request!",
		})
	})
	router.POST("/", func(c *gin.Context) {

	})

	// Auth route - Grouping route
	auth := router.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
		auth.POST("/logout", controller.Logout)
	}

	// API route - Grouping route
	api := router.Group("/api", middleware.Authentication())
	{
		// User route
		api.GET("/users", controller.GetUsers)
		user := api.Group("/user")
		{
			user.GET("/:id", controller.GetUserById)
			user.DELETE("/:id", controller.DeleteUserById)
			user.PUT("/:id", controller.UpdateUserById)
			user.POST("/", controller.CreateUser)
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
			// category.POST("/", middleware.Authentication(), controller.CreateCategory)
			category.POST("/", controller.CreateCategory)
			category.GET("/:id", controller.GetCategoryById)
			category.DELETE("/:id", controller.DeleteCategoryById)
			category.PUT("/:id", controller.UpdateCategoryById)
			category.GET("/:id/recipes", controller.GetCategoryRecipes)
		}

		// Article route
		api.GET("/articles", controller.GetArticles)
		article := api.Group("/article")
		{
			article.GET("/:id", controller.GetArticleById)
			article.DELETE("/:id", controller.DeleteArticleById)
			article.PUT("/:id", controller.UpdateArticleById)
			// article.POST("/", middleware.Authentication(), controller.CreateArticle)
			article.POST("/", controller.CreateArticle)
			article.GET("/categories", controller.GetArticleCategories)
		}

		// Bookmark route
		bookmark := api.Group("/bookmark")
		{
			bookmark.GET("/:id", controller.GetBookmarkById)
			bookmark.DELETE("/:id", controller.DeleteBookmarkById)
			bookmark.PUT("/:id", controller.UpdateBookmarkById)
			bookmark.POST("/", controller.CreateBookmark)
		}
	}

	return router
}
