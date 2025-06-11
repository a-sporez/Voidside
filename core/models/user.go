// /models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"unique" json:"name"`
	Email string `gorm:"unique" json:"email"`
}
