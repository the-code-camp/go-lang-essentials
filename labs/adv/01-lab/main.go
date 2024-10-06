package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	GPA  float64
}

type StudentManager struct {
	Students []Student
}

func (sm StudentManager) AddStudent(s Student) {
	// Implement the logic to add a student to the list.
}

func (sm StudentManager) DisplayStudents() {
	// Implement the logic to display all students and their details.
}

func (sm StudentManager) UpdateStudent(name string, updatedStudent Student) {
	// Implement the logic to update a student's details by name.
}

func (sm StudentManager) DeleteStudent(name string) {
	// Implement the logic to delete a student by name.
}

func main() {
	sm := StudentManager{} // Create a new StudentManager instance

	// Example usage:
	// Add students
	sm.AddStudent(Student{Name: "Alice", Age: 20, GPA: 3.5})
	sm.AddStudent(Student{Name: "Bob", Age: 21, GPA: 3.2})
	sm.AddStudent(Student{Name: "Charlie", Age: 22, GPA: 3.9})

	// Display students
	fmt.Println("Students:")
	sm.DisplayStudents()

	// Update student
	sm.UpdateStudent("Bob", Student{Name: "Bob", Age: 22, GPA: 3.3})
	fmt.Println("\nAfter updating Bob's details:")
	sm.DisplayStudents()

	// Delete student
	sm.DeleteStudent("Charlie")
	fmt.Println("\nAfter deleting Charlie:")
	sm.DisplayStudents()
}
