package json

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func E404(c echo.Context) error {
	return c.JSON(http.StatusNotFound, struct {
		Message string `json:"message"`
	}{"Not found"})
}
