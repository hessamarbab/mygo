package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"sika-hessam/database"
	"sika-hessam/router"
	"sika-hessam/seeders"
)

func main() {
	initEnv()
	//init db
	defer database.CloseConnection()
	database.ConnectToPgrs()
	database.AutoMigrate()
	seeders.SeedUsers() // todo uncomment

	//http app
	e := echo.New()

	router.Router(e)

	e.Logger.Fatal(e.Start(":80"))
}

func initEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
