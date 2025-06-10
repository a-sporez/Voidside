// /controller/posts.go
// Each function here is a handler for an API route
package controllers

import (
	"net/http"

	"core/config"
	"core/models"

	"github.com/gin-gonic/gin"
)

// Read all posts in the database
func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

// Read one single post by id or 404 if not found
func GetPost(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot find post"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// Write new post from JSON input
func CreatePost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Title: input.Title, Content: input.Content}
	config.DB.Create(&post)
	c.JSON(http.StatusCreated, post)
}

// Modifies a post based on JSON inpit and ID
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot find post"})
		return
	}
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&post).Updates(input)
	c.JSON(http.StatusOK, post)
}

// Removes post by ID
func DeletePost(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot find post"})
		return
	}
	config.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
