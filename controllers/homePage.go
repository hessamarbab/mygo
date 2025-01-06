package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetHome(c echo.Context) error {
	return c.HTML(http.StatusOK,
		"Server is running; you could use apis!")
}
