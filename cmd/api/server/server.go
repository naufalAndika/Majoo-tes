package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// New instance of server
func New() *echo.Echo {
	e := echo.New()
	e.GET("/", healthCheck)

	return e
}

// Start new server
func Start(e *echo.Echo) {
	e.Start(":8080")
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status": "OK",
	})
}
