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

type UserController struct {
	userRepo repositories.UserRepo
}

func NewUserController(userRepo repositories.UserRepo) *UserController {
	return &UserController{userRepo: userRepo}
}

func (cntlr *UserController) GetUser(c echo.Context) error {

	var reqBody GetUserRequestBody

	err := c.Bind(&reqBody)
	if err != nil || reqBody.Id == "" {
		return json.E400(c, "")
	} else {

		err, user := cntlr.userRepo.GetUserWithAddressesById(reqBody.Id)
		if err != nil {
			return json.E404(c, "")
		}

		return c.JSON(http.StatusOK, user)
	}
}
