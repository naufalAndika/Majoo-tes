package middleware

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/naufalAndika/Majoo-tes/cmd/api/request"
)

// UserMiddleware for checking user token
func UserMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := request.ParseToken(c)
			if err != nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			attachTokenToContext(token, c)

			return next(c)
		}
	}
}

func attachTokenToContext(token *request.TokenPayload, c echo.Context) {
	c.Set("user_id", token.UserID)
	c.Set("is_authenticated", true)
}
