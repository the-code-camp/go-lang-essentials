package model

type Customer struct {
	Id          string `json:"id" db:"customer_id"`
	Name        string
	City        string `json:"city"`
	Zipcode     string
	DateofBirth string `json:"-" db:"date_of_birth"`
	Status      string
}
