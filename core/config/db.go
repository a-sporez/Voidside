// /config/db.go
package config

import (
	"log"

	"core/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// used globally so other packages can access db
var DB *gorm.DB

// gorm.Open() connects to SQLite DB
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db:", err)
	}
	database.AutoMigrate(&models.Post{}) // generate tables
	DB = database                        // Assign to global
}
