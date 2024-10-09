package models

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("Firstname: %s, Lastname: %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

type Employee struct {
	Person
	EmployeeID string
}
