package json

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func E404(c echo.Context, message string) error {
	if message == "" {
		message = "Not found"
	}
	return c.JSON(http.StatusNotFound, struct {
		Message string `json:"message"`
	}{message})
}
