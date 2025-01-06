package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"sika-hessam/controllers"
	"sika-hessam/database"
	"sika-hessam/repositories"
	"sika-hessam/router"
	"sika-hessam/seeders"
)

func main() {
	initEnv()
	//init db
	defer database.CloseConnection()
	database.ConnectToPgrs()
	database.AutoMigrate()
	seeders.SeedUsers()

	//http app
	e := echo.New()
	userController := controllers.NewUserController(&repositories.UserPgsRepo{})

	router.Router(e, userController)

	e.Logger.Fatal(e.Start(":80"))
}

func initEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
