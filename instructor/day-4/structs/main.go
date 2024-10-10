package main

import (
	"fmt"

	"github.com/the-code-camp/go-lang-essentials/models"
)

func main() {
	// mutation example
	p := models.Person{Salary: 1000, FirstName: "firstname"}
	fmt.Println("Salary Before: ", p.Salary)
	p.AwardBonus(1000)
	(&p).AwardBonus(1000)
	fmt.Println("Salary After: ", p.Salary)
}

func main_example() {
	// var m map[string]models.Person

	m := make(map[string]models.Person)

	var p models.Person
	p.FirstName = "some-firstname"
	p.LastName = "lastname"
	p.Age = 20

	m[p.FirstName] = p

	fmt.Println(m)
	// err := p.ValidatePerson()

	// if err != nil {
	// 	fmt.Println("Some error: ", err)
	// }
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
