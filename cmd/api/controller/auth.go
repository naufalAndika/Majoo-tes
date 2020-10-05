package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naufalAndika/Majoo-tes/cmd/api/request"
	"github.com/naufalAndika/Majoo-tes/internal/service"
)

// Auth UserController model
type Auth struct {
	service *service.Auth
}

// NewAuth user request handler
func NewAuth(svc *service.Auth, ar *echo.Group) {
	ac := Auth{svc}

	ar.POST("/login", ac.login)
}

func (ac *Auth) login(c echo.Context) error {
	req, err := request.ParseLogin(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse request body")
	}

	userAuth, err := ac.service.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, userAuth)
}
