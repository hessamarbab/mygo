package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sika-hessam/database"
	"sika-hessam/models"
)

type GetUserRequestBody struct {
	Id string `json:"id"`
}

func GetUser(c echo.Context) error {
	db := database.GetDB()
	var reqBody GetUserRequestBody
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{"Bad request"})
	} else {
		var user models.User
		user.Id = reqBody.Id
		err := db.Model(&models.User{}).Preload("Addresses").Find(&user, "id = ?", reqBody.Id).Error

		if err != nil {
			return c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
			}{"Not found"})
		}
		return c.JSON(http.StatusOK, user)
	}
}
