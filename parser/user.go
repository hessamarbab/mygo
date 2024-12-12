package parser

import (
	"encoding/json"
	"sika-hessam/dto"
	"strings"
)

type userFormat struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Addresses []dto.Address `json:"addresses"`
}

func UsersParse(dat []byte) ([]dto.User, error) {
	ufs := make([]userFormat, 0)
	if err := json.Unmarshal(dat, &ufs); err != nil {
		return nil, err
	}
	var users []dto.User

	for _, uf := range ufs {
		name := strings.Split(uf.Name, " ")
		users = append(users, dto.User{
			Id:        uf.Id,
			FirstName: name[0],
			LastName:  name[1],
			Addresses: uf.Addresses,
		})
	}
	return users, nil
}
