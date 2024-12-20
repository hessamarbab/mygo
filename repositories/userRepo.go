package repositories

import (
	"sika-hessam/database"
	"sika-hessam/models"
)

func GetUserWithAddressesById(id string) (error, *models.User) {
	db := database.GetDB()
	var user models.User
	if err := db.Model(&models.User{}).Preload("Addresses").First(&user, "id = ?", id).Error; err != nil {
		return err, nil
	}
	return nil, &user
}
