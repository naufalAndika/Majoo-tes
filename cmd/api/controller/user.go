package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naufalAndika/Majoo-tes/internal/service"
)

// User UserController model
type User struct {
	service *service.User
}

// NewUser user request handler
func NewUser(svc *service.User, ur *echo.Group) {
	uc := User{svc}

	ur.GET("/", uc.findAllUsers)
}

func (uc *User) findAllUsers(c echo.Context) error {
	users, err := uc.service.FindAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
