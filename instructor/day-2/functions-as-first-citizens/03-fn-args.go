package main

import "fmt"

func main() {
	exec(f1)
	exec(f2)

	exec(func() {
		fmt.Println("anonymous function passed to execute")
	})
	// we use this concept in function composition
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3(x, y int) {
	fmt.Println("f2 invoked")
}
