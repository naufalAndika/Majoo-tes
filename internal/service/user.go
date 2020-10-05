package service

import (
	model "github.com/naufalAndika/Majoo-tes/internal"
)

// User service model
type User struct {
	udb model.UserDB
}

// NewUser create new User Service instance
func NewUser(udb model.UserDB) *User {
	return &User{udb}
}

// FindAllUsers in database
func (u *User) FindAllUsers() (*[]model.User, error) {
	return u.udb.FindAllUsers()
}
