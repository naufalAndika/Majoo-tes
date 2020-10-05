package request

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ID get id from param
func ID(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest)
	}

	return id, nil
}
