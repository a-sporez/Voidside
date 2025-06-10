// /config/db.go
package config

import (
	"log"
	"os"

	"core/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// used globally so other packages can access db
var DB *gorm.DB

// gorm.Open() connects to SQLite DB
func ConnectDatabase() {
	// load DB path from env
	_ = godotenv.Load("env/.env")
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "default.db" // fallback
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	database.AutoMigrate(&models.Post{}) // generate tables
	DB = database                        // Assign to global
}
