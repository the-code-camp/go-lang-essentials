package main

import "fmt"

func main() {
	var name string
	var counter int
	fmt.Printf("%s, %v\n", name, counter)

	var flag bool
	fmt.Println(flag)

	// Declare variables without initializing (they take zero values)
	var i int            // Zero value for int
	var f float64        // Zero value for float64
	var s string         // Zero value for string
	var b bool           // Zero value for bool
	var p *int           // Zero value for pointer (nil)
	var arr []int        // Zero value for slice (nil)
	var m map[string]int // Zero value for map (nil)

	// Print zero values
	fmt.Println("Zero value of int:", i)
	fmt.Println("Zero value of float64:", f)
	fmt.Println("Zero value of string:", s)
	fmt.Println("Zero value of bool:", b)
	fmt.Println("Zero value of pointer:", p)
	fmt.Println("Zero value of slice:", arr)
	fmt.Println("Zero value of map:", m)

	// Declare and print struct
	type Person struct {
		name string
		age  int
	}

	var person Person // Zero value for struct (fields are zero values)
	fmt.Println("Zero value of struct:", person)
}
