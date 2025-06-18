package models

type ChatMessage struct {
	UserID  string `json:"userId"`
	Message string `json:"message"`
}
