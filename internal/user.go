package model

import (
	"time"
)

// User model
type User struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Username  string     `gorm:"size:255;not null;unique" json:"username"`
	Password  string     `gorm:"size:255;not null;" json:"password"`
	Fullname  string     `gorm:"size:255;not null;" json:"fullname"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// UserDB interface
type UserDB interface {
	FindAllUsers() (*[]User, error)
}
