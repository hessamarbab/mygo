package models

import (
	"gorm.io/gorm"
	"sika-hessam/dto"
)

type Address struct {
	Street   string `gorm:"size:32"`
	City     string `gorm:"size:32"`
	State    string `gorm:"size:32"`
	Zip_code string `gorm:"size:32"`
	Country  string `gorm:"size:64"`
	gorm.Model
	UserId string
	User   User `gorm:"foreignKey:UserId"`
}

func (address *Address) Fill(addrDara dto.Address) {
	address.Street = addrDara.Street
	address.City = addrDara.City
	address.State = addrDara.State
	address.Zip_code = addrDara.Zip_code
	address.Country = addrDara.Country
}
