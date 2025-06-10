// /controllers/users.go

package controllers

import (
	"net/http"

	"core/config"
	"core/models"

	"github.com/gin-gonic/gin"
)

// Writes new entry to the databasa.
func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Name: input.Name, Email: input.Email}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Return all users from database.
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// Return single user by ID from database.
func GetUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Replace single user by ID with new parameters from input.
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot find user"})
		return
	}
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user)
}

// Remove single user completely from database.
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.First(&user, c.Param("id")).Error; err != nil {
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
