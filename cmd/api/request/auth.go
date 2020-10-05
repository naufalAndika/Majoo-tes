package request

import "github.com/labstack/echo"

// Login request class for login
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// ParseLogin parse request body for login
func ParseLogin(c echo.Context) (*Login, error) {
	req := new(Login)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return req, nil
}
