package main

import "fmt"

func main() {
	var i int = 10
	var name = "golang"

	var c rune = 'A'

	var flag = false

	fmt.Println(i, name, c, flag)

	// array is fixed length
	// slice is dynamic array, is supported by array [in JAVA as ArrayList]

	var city string
	var city1 = "New Delhi"

	city = "Chicago"

	city2 := "New Delhi"
	num := 10
	city2 = "New York"

	fmt.Println(city, city1, city2, num)

}
