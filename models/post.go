// /models/post.go
package models

import "gorm.io/gorm"

// gorm.Model() includes ID, CreatedAy, UpdatedAt, DeletedAt
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
