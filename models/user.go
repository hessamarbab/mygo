package models

import (
	"gorm.io/gorm"
	"sika-hessam/dto"
	"time"
)

type User struct {
	Id        string         `gorm:"primaryKey"`
	Firstname string         `gorm:"size:32"`
	Lastname  string         `gorm:"size:32"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Addresses []Address
}

func (user *User) Fill(usrData dto.User) {
	user.Id = usrData.Id
	user.Firstname = usrData.FirstName
	user.Lastname = usrData.LastName
	for _, addr := range usrData.Addresses {
		var address Address
		address.Fill(addr)
		user.Addresses = append(user.Addresses, address)
	}
}
