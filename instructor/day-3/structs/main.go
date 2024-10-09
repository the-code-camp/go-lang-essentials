package main

import (
	"fmt"

	"github.com/the-code-camp/go-lang-essentials/models"
)

func main() {
	var p models.Person
	p.FirstName = ""

	// err := models.ValidatePerson(p)

	err := p.ValidatePerson()

	if err != nil {
		fmt.Println("Some error: ", err)
	}
}

func personsExample() {
	var p models.Person

	p.FirstName = "firstname"
	p.LastName = "lastname"
	p.Age = 10

	p1 := models.Person{
		Age:       20,
		FirstName: "some-first",
	}

	var e1 models.Employee

	e1.FirstName = "emp-firstname"
	e1.LastName = "emp-lastname"
	e1.EmployeeID = "empid"

	e := models.Employee{
		Person: models.Person{
			FirstName: "",
		},
		EmployeeID: "123",
	}

	fmt.Println(p)
	fmt.Println(p1)

	fmt.Println(e)
	fmt.Println(e1)
}
