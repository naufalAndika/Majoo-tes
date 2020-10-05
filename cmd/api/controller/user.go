package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naufalAndika/Majoo-tes/cmd/api/request"
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
	ur.GET("/:id/", uc.findByID)
}

func (uc *User) findAllUsers(c echo.Context) error {
	users, err := uc.service.FindAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (uc *User) findByID(c echo.Context) error {
	userID, err := request.ID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to fetch ID")
	}

	user, err := uc.service.FindByID(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
