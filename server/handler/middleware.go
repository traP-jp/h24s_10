package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const traQIDKey = "traq-id"

// TraQIDMiddleware is a middleware for traQ ID. It sets the traQ ID from the X-Forwarded-User header to the context.
// If the header is not set, it returns 401 Unauthorized.
func TraQIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traqID := c.Request().Header.Get("X-Forwarded-User")
		if traqID == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "X-Forwarded-User header is required")
		}
		c.Set(traQIDKey, traqID)
		return next(c)
	}
}

// DevTraQIDMiddleware is a middleware for development. It sets a fixed traQ ID "dev-traq-id" to the context.
func DevTraQIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(traQIDKey, "Luftalian")
		return next(c)
	}
}
