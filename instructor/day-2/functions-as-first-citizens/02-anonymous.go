package main

import "fmt"

func main() {

	func() {
		fmt.Println("anonymous function invoked")
	}()

	func(x, y int) {
		result := x + y
		fmt.Println("Result is: ", result)
	}(100, 200)

	fmt.Println(func(x, y int) int {
		return x + y
	}(100, 200))
}
