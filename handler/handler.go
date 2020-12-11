package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// RootHandler returns a simple response to the request against the root URI.
func RootHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
