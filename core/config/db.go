// /config/db.go
package config

import (
	"log"
	"os"

	"core/models"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// used globally so other packages can access db
var DB *gorm.DB

// gorm.Open() connects to SQLite DB
func ConnectDatabase() {
	// load DB path from env
	_ = godotenv.Load()
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "default.db" // fallback
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	// generate tables and assign to global database
	database.AutoMigrate(&models.Post{}, &models.User{})
	DB = database
}
