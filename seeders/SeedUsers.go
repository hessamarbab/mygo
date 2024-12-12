package seeders

import (
	"log"
	"os"
	"sika-hessam/database"
	"sika-hessam/dto"
	"sika-hessam/models"
	"sika-hessam/parser"
)

func SeedUsers() {
	users := readFile()
	ch := make(chan dto.User, 10)
	quit := make(chan bool)
	for i := 0; i < 10; i++ {
		go insertToDB(ch, quit)
	}
	for _, usr := range users {
		ch <- usr
	}
	for i := 0; i < 10; i++ {
		quit <- false
	}
	close(ch)
	close(quit)
	log.Println("seed Done!")
}

func insertToDB(ch chan dto.User, quit chan bool) {
	for {
		select {
		case usr := <-ch:
			var user models.User
			user.Fill(usr)
			dbClient := database.GetDB()
			err := dbClient.Create(&user).Error
			if err != nil {
				log.Fatal("Error creating users")
			}
		case <-quit:
			return
		}
	}
}

func readFile() []dto.User {
	dat, err := os.ReadFile("./users_data.json")
	if err != nil {
		log.Fatal("Error read file  ./users_data.json")
	}
	users, err := parser.UsersParse(dat)
	if err != nil {
		log.Fatal(err)
	}
	return users
}
