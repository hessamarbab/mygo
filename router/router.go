package router

import (
	"github.com/labstack/echo/v4"
	"sika-hessam/controllers"
)

func Router(e *echo.Echo) {

	e.GET("/", controllers.GetHome)
	e.GET("/user", controllers.GetUser)
}
