package main

import "fmt"

func main() {
	var fn func() // what is the type of function?
	// functions are type safe in Go

	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var addProcessor func(x, y int)

	addProcessor = func(x, y int) {
		result := x + y
		fmt.Println("Result is: ", result)
	}
	addProcessor(100, 200)

	mathOperation := func(x, y int) int {
		return x + y
	}
	fmt.Println(mathOperation(100, 200))

	mathOperation = func(x, y int) int {
		return x - y
	}
	fmt.Println(mathOperation(100, 200))

	// understanding how you define the function type is very important, it helps you
	// with two concepts
	// - passing functions as arguments
	// - returning a function as return value
}
