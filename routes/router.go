// /routes/router.go
package routes

import (
	"Voidside/controllers"

	"github.com/gin-gonic/gin"
)

// Maps endpoints to controllers and returns the Gin engine.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/posts", controllers.GetPost)
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	return r
}
