package models

import (
	"errors"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Salary    float64
	// cannot do things like this
	// func Validate() error {

	// }
}

func (p *Person) AwardBonus(bonus float64) {
	p.Salary += bonus
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// this becomes a method from a function
// method is associated with Person struct
// receiver functions in Go
func (p Person) ValidatePerson() error {
	if p.FirstName == "" {
		return errors.New("first name cannot be blank")
	}

	if p.LastName == "" {
		return errors.New("last name cannot be blank")
	}

	return nil
}

func (p Person) String() string {
	return fmt.Sprintf("Firstname: %s, Lastname: %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

type Employee struct {
	Person
	EmployeeID string
}
