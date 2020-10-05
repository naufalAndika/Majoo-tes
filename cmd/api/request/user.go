package request

import "github.com/labstack/echo"

// CreateUser request class for create user
type CreateUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
}

// ParseCreateUser parse request body for create user
func ParseCreateUser(c echo.Context) (*CreateUser, error) {
	req := new(CreateUser)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return req, nil
}
