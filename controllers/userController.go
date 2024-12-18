package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"sika-hessam/database"
	"sika-hessam/models"
)

func GetUser(c echo.Context) error {
	db := database.GetDB()
	json_map := make(map[string]string)
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{"Bad request"})
	} else {
		var user models.User
		user.Id = json_map["id"]
		err := db.Model(&models.User{}).Preload("CreditCards").Find(&user, "id = ?", json_map["id"]).Error

		if err != nil {
			return c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
			}{"Not found"})
		}
		return c.JSON(http.StatusOK, user)
	}
}
