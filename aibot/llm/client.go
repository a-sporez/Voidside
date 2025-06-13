package llm

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"aibot/internal"
	"aibot/models"
)

func CallLLM(userID string) (string, error) {
	llmURL := os.Getenv("LLM_URL")
	bearerToken := os.Getenv("LLM_TOKEN") // this is safer than hardcoding
	/*
		if llmURL == "" || bearerToken == "" {
			log.Fatal("Missing LLM_URL or LLM_TOKEN in .env")
		}
	*/
	history := internal.TrimMemory(internal.MemoryStore[userID], 5)
	payload := models.LLMRequest{
		Messages:    history, // use context window memory
		Temperature: 0.7,
		Stream:      false,
	}

	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", llmURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Printf("LLM response body: %s", body)

	var result models.LLMResponse
	if err := json.Unmarshal(body, &result); err != nil {
		// check for silent failures
		log.Printf("Unmarshall error: %v", err)
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	if len(result.Choices) == 0 {
		log.Printf("No choices returned by LLM. Full body: %s", body)
	}

	log.Printf("No choices returned by LLM")
	return "{no reply}", nil
}
