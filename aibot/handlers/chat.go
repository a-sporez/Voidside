package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"aibot/internal"
	"aibot/llm"
	"aibot/models"
)

func HandleChat(c *gin.Context) {
	var input models.ChatInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// validate input
	if input.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	now := time.Now() // rate limiter 3s cooldown
	if last, ok := internal.LastSeen[input.UserID]; ok && now.Sub(last) < 3*time.Second {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded."})
		return
	}
	internal.LastSeen[input.UserID] = now

	// store user message in memory and ignore blank strings
	trimmed := strings.TrimSpace(input.Message)
	if trimmed != "" {
		internal.MemoryStore[input.UserID] = append(
			internal.MemoryStore[input.UserID], models.Message{
				Role:    "user",
				Content: input.Message,
			})
	}

	reply, err := llm.CallLLM(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LLM error"})
		return
	}

	// store bot reply in memory after trimming blank strings
	if trimmed := strings.TrimSpace(reply); trimmed != "" {
		internal.MemoryStore[input.UserID] = append(internal.MemoryStore[input.UserID], models.Message{
			Role:    "assistant",
			Content: reply,
		})
	}

	c.JSON(http.StatusOK, models.ChatOutput{Reply: reply})
}
