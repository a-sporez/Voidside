// /main.go
package main

import (
	"core/config"
	"core/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	config.ConnectDatabase() // instantiate DB and store globally

	//middleware.InitJWT() // initialize keycloak JWKS

	r := routes.SetupRouter() // Return Gin engine with routes mounted.

	port := os.Getenv("PORT") // get port from .env
	if port == "" {
		port = "8080"
	}
	log.Println("Running on port:", port)

	r.Run(":" + port) // start server on default port
}
