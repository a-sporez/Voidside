// /routes/router.go
package routes

import (
	"core/controllers"

	"github.com/gin-gonic/gin"
)

// Maps endpoints to controllers and returns the Gin engine.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPost)
	router.POST("/posts", controllers.CreatePost)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	return router
}
