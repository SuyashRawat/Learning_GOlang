package models

import "gorm.io/gorm"

// User represents the user model for the database
type User struct {
	gorm.Model
	FirstName    string `gorm:"size:100;not null" json:"first_name"` // JSON tag added
	LastName     string `gorm:"size:100;not null" json:"last_name"`  // JSON tag added
	Email        string `gorm:"unique;not null" json:"email"`        // JSON tag added
	Phone        string `gorm:"size:15;not null" json:"phone"`       // JSON tag added
	Token        string `gorm:"size:255" json:"token"`               // JSON tag added
	UserType     string `gorm:"size:50;not null" json:"user_type"`   // JSON tag added
	RefreshToken string `gorm:"size:255" json:"refresh_token"`       // JSON tag added
	Password     string `gorm:"size:255;not null" json:"password"`   // JSON tag added
}
