package models

// incoming message structure
type ChatInput struct {
	UserID  string `json:"userId"` // from discord or internal systems
	Message string `json:"message"`
}

// outgoing reply structure
type ChatOutput struct {
	Reply string `json:"reply"`
}

type LLMRequest struct {
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
	Stream      bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
