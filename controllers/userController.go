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

	err := c.Bind(&reqBody)
	if err != nil || reqBody.Id == "" {
		return json.E400(c, "")
	} else {

		err, user := repositories.GetUserWithAddressesById(reqBody.Id)
		if err != nil {
			return json.E404(c, "")
		}

		return c.JSON(http.StatusOK, user)
	}
}
