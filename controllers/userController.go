package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sika-hessam/errors/json"
	"sika-hessam/repositories"
)

type GetUserRequestBody struct {
	Id string `json:"id"`
}

func GetUser(c echo.Context) error {

	var reqBody GetUserRequestBody

	if err := c.Bind(&reqBody); err != nil {
		return json.E400(c)
	} else {

		err, user := repositories.GetUserWithAddressesById(reqBody.Id)
		if err != nil {
			return json.E404(c)
		}

		return c.JSON(http.StatusOK, user)
	}
}
