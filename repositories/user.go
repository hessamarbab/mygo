package repositories

import (
	"fmt"
	"sika-hessam/database"
	"sika-hessam/models"
)

func GetUserWithAddressesById(id string) (error, *models.User) {
	db := database.GetDB()
	var user models.User
	if err := db.Model(&models.User{}).Preload("Addresses").First(&user, "id = ?", id).Error; err != nil {
		fmt.Println("jojo1")
		return err, nil
	}
	fmt.Println("jojo4")
	return nil, &user
}
