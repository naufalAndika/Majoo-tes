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

// CreateUser store new user
func (u *User) CreateUser(username, password, fullname string) (*model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
		Fullname: fullname,
	}

	return u.udb.CreateUser(user)
}

// FindAllUsers in database
func (u *User) FindAllUsers() (*[]model.User, error) {
	return u.udb.FindAllUsers()
}

// FindByID find user by id
func (u *User) FindByID(ID uint) (*model.User, error) {
	return u.udb.FindByID(ID)
}

// DeleteByID delete user by id
func (u *User) DeleteByID(ID uint) error {
	return u.udb.DeleteByID(ID)
}
