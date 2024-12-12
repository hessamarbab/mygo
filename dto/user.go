package dto

type User struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Addresses []Address `json:"addresses"`
}
