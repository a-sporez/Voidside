package models

type ChatRequest struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"message"`
}
