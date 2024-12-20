package json

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func E400(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, struct {
		Message string `json:"message"`
	}{"Bad request"})
}
