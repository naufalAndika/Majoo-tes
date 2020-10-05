package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	model "github.com/naufalAndika/Majoo-tes/internal"
)

// Auth service model
type Auth struct {
	udb model.UserDB
}

// UserAuthentication authenticate user response
type UserAuthentication struct {
	User  *model.User `json:"user,omitempty"`
	Token string      `json:"token"`
}

// NewAuth create new User Service instance
func NewAuth(udb model.UserDB) *Auth {
	return &Auth{udb}
}

// Login return user and token when credential is valid
func (a *Auth) Login(username, password string) (*UserAuthentication, error) {
	user, err := a.udb.FindByCredentials(username, password)
	if err != nil {
		return nil, err
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	return &UserAuthentication{
		user,
		token,
	}, nil
}

func generateToken(user *model.User) (string, error) {
	expire := time.Now().Add(24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expire,
	})

	return token.SignedString([]byte("secret"))
}
