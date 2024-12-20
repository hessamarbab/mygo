package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sika-hessam/repositories"
)

type GetUserRequestBody struct {
	Id string `json:"id"`
}

func GetUser(c echo.Context) error {
	var reqBody GetUserRequestBody
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{"Bad request"})
	} else {
		err, user := repositories.GetUserWithAddressesById(reqBody.Id)
		if err != nil {
			return c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
			}{"Not found"})
		}
		return c.JSON(http.StatusOK, user)

	}
}
