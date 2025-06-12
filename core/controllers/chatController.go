package controllers

import (
	"core/dto"
	"core/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Accepts a chat message and relays it to the aibot.
func ProxyChatHandler(c *gin.Context) {
	tokenRaw, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}
	claims := tokenRaw.(*jwt.Token).Claims.(jwt.MapClaims)

	userID := claims["preferred_username"].(string)

	var input dto.ChatRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat input"})
		return
	}

	reply, err := services.SendToAIBot(userID, input.Message)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "airbot error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ChatResponse{Reply: reply})
}
