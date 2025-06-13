// SporeDrop/main.go
package main

import (
	"aibot/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found; using system env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("chatbot API running on :" + port)

	router := gin.Default()
	router.POST("/chat", handlers.HandleChat)
	router.Run(":" + port)
}
