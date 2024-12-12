package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetHome(c echo.Context) error {
	return c.HTML(http.StatusOK,
		"<html><h1>Hi use /user to get user by id like this:</h1>"+
			"<h3><pre>           (use method GET & this headers "+
			"Accept: \"application/json\" ,  Content-Type:\"application/json\" )\n\n"+
			"			{\n"+
			"				\"id\": \"35b75a6b-0183-43e1-bdd0-3cd062dc6935\"\n"+
			"			}</pre></h3>"+
			"</html>")
}
