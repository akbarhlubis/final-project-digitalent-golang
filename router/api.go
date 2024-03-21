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

		// User route
		api.POST("/users", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Get all users",
			})
		})
	}

	return router
}
