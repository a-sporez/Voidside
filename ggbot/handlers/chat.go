package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"ggbot/models"

	"github.com/bwmarrin/discordgo"
)

func HandleChat(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore bot messages.
	if m.Author.Bot {
		return
	}

	// Basic command trigger.
	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	message := strings.TrimPrefix(m.Content, "!")
	payload := models.ChatRequest{
		User:    m.Author.Username,
		Message: message,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", os.Getenv("API_URL"), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer"+os.Getenv("API_SECRET"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		s.ChannelMessageSend(m.ChannelID, "API Error")
		return
	}
	defer resp.Body.Close()

	var output models.ChatResponse
	json.NewDecoder(resp.Body).Decode(&output)
	s.ChannelMessageSend(m.ChannelID, output.Reply)
}
