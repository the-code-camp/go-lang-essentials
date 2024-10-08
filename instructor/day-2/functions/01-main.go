package main

import "fmt"

func main() {

	result := sum(10, 20)
	fmt.Println("Result is: ", result)

}

func sum(x, y int) (result int) { // you can also make the return value as a named argument
	result = x + y
	return result
}

func Write(b []byte) (int, error) {
	return 0, nil
}
