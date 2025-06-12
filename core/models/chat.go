package models

type ChatMessage struct {
	UserID  string `json:"user"`
	Message string `json:"message"`
}
