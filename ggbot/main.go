package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ggbot/handlers"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using system env")
	}

	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN not set in .env")
	}

	session, err := discordgo.New(token)
	if err != nil {
		log.Fatal("Failed to create discord session:", err)
	}

	session.AddHandler(handlers.HandleChat)

	err = session.Open()
	if err != nil {
		log.Fatal("Cannot open connection to Discord:", err)
	}
	log.Println("ggbot initialized")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down ggbot...")
	session.Close()
}
