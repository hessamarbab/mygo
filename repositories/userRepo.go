package repositories

import (
	"sika-hessam/database"
	"sika-hessam/models"
)

type UserRepo interface {
	GetUserWithAddressesById(id string) (error, *models.User)
}

type UserPgsRepo struct {
}

func (repo *UserPgsRepo) GetUserWithAddressesById(id string) (error, *models.User) {
	db := database.GetDB()
	var user models.User
	if err := db.Model(&models.User{}).Preload("Addresses").First(&user, "id = ?", id).Error; err != nil {
		return err, nil
	}
	return nil, &user
}
