package dto

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip_code string `json:"zip_Code"`
	Country  string `json:"country"`
}
