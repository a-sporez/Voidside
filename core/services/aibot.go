package services

import (
	"bytes"
	"core/dto"
	"core/models"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// Sends user message to aibot and returns response.
func SendToAIBot(userID string, message string) (string, error) {
	msg := models.ChatMessage{
		UserID:  userID,
		Message: message,
	}

	body, _ := json.Marshal(msg)
	aibotURL := os.Getenv("AIBOT_URL")
	if aibotURL == "" {
		return "", errors.New("AIBOT_URL not set in env")
	}

	req, err := http.NewRequest("POST", aibotURL+"/chat", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var output dto.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return "", err
	}

	return output.Reply, nil
}
