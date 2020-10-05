package request

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login request class for login
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TokenPayload auth token payload
type TokenPayload struct {
	UserID uint
}

// ParseLogin parse request body for login
func ParseLogin(c echo.Context) (*Login, error) {
	req := new(Login)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return req, nil
}

// ParseToken parse auth token
func ParseToken(c echo.Context) (*TokenPayload, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Token not supplied")
	}

	return getTokenPayload(token)
}

func getTokenPayload(token string) (*TokenPayload, error) {
	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	if !payload.Valid {
		return nil, err
	}

	claims := payload.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, err
	}

	return &TokenPayload{
		UserID: uint(userID),
	}, nil
}
