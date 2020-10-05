package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naufalAndika/Majoo-tes/cmd/api/middleware"
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

	ur.POST("/", uc.createUser)
	ur.GET("/", uc.findAllUsers)
	ur.GET("/:id", uc.findByID)
	ur.GET("/me", uc.userDetail, middleware.UserMiddleware())
	ur.DELETE("/:id", uc.deleteByID)
	ur.DELETE("/me", uc.deleteCurrentUser, middleware.UserMiddleware())
	ur.PUT("/:id", uc.updateByID)
	ur.PUT("/me", uc.updateCurrentUser, middleware.UserMiddleware())
}

func (uc *User) createUser(c echo.Context) error {
	req, err := request.ParseCreateUser(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse request body")
	}

	user, err := uc.service.CreateUser(req.Username, req.Password, req.Fullname)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
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

func (uc *User) userDetail(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user")
	}

	user, err := uc.service.FindByID(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *User) deleteByID(c echo.Context) error {
	userID, err := request.ID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to fetch ID")
	}

	if err := uc.service.DeleteByID(uint(userID)); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "User deleted")
}

func (uc *User) deleteCurrentUser(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user")
	}

	if err := uc.service.DeleteByID(userID); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "User deleted")
}

func (uc *User) updateByID(c echo.Context) error {
	userID, err := request.ID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to fetch ID")
	}

	req, err := request.ParseUpdateUser(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse request body")
	}

	user, err := uc.service.UpdateByID(uint(userID), req.Password, req.Fullname)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *User) updateCurrentUser(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch user")
	}

	req, err := request.ParseUpdateUser(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to parse request body")
	}

	user, err := uc.service.UpdateByID(userID, req.Password, req.Fullname)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
