// /routes/router.go
package routes

import (
	"core/controllers"
	"core/middleware"

	"github.com/gin-gonic/gin"
)

// Maps endpoints to controllers and returns the Gin engine.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// public routes
	router.POST("/users", controllers.CreateUser)

	router.POST("/proxy/chat", controllers.ProxyChatHandler)

	// protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/posts", controllers.GetPosts)
		protected.GET("/posts/:id", controllers.GetPost)
		protected.POST("/posts", controllers.CreatePost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)

		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUser)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
	}

	return router
}
