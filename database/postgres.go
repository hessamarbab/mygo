package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sika-hessam/models"
)

var dbClient *gorm.DB

func ConnectToPgrs() {
	var err error
	dsn := os.Getenv("DB_DSN")
	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database (init)")
	}
	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal("failed to connect database (ping)")
	}

	log.Println("Database Connected")
}
func AutoMigrate() {
	if err := dbClient.AutoMigrate(models.User{}); err != nil {
		log.Fatal(err)
	}
	if err := dbClient.AutoMigrate(models.Address{}); err != nil {
		log.Fatal(err)
	}
}
func GetDB() *gorm.DB {
	return dbClient
}
func CloseConnection() {
	sqlDB, _ := dbClient.DB()
	err := sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}
}
