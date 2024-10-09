package main

import "fmt"

func main() {

	// pointers are type safe, this is also one of the difference between C, C++ vs Go
	// var noPtr *int // what should be the type?
	// noPtr = &no

	// fmt.Println(no == *(&no)) // should be same?

	// var input int
	// fmt.Scanln(&input)

	no := 10
	fmt.Println("Before incrementing: ", no)
	increment(&no)
	fmt.Println("After incrementing: ", no)
}

func increment(x *int) {
	// fmt.Println("Pointer address: ", x, *x)
	// *x = 200
	*x++
}
