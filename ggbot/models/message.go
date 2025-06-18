package models

type ChatRequest struct {
	User    string `json:"userId"`
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"message"`
}
