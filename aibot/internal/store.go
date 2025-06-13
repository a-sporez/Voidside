package internal

import (
	"aibot/models"
	"time"
)

// userID -> last request timestamp.
var LastSeen = make(map[string]time.Time)

// Create context memory map.
var MemoryStore = make(map[string][]models.Message)

// Defines the limit of cached memory.
func TrimMemory(messages []models.Message, limit int) []models.Message {
	if len(messages) <= limit {
		return messages
	}
	return messages[len(messages)-limit:]
}
