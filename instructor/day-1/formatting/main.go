package main

import (
	"fmt"
)

func main() {
	p := "New Delhi"
	fmt.Printf("City name is: %q\n", p)
	fmt.Printf("type: %T\n", p)

	charA, charZ := 97, 122
	fmt.Printf("type: %c and %c\n", charA, charZ)
	fmt.Printf("type: %d and %d\n", charA, charZ)

	fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("width1: |%6d|%6d|\n", 19, 731)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 6.7, 12.82)

	fmt.Printf("%6.2f|\n", 6.7)
	fmt.Printf("%6.2f|\n", 12.82)
	fmt.Println("----------")

	fmt.Printf("%-6.2f|\n", 6.7)
	fmt.Printf("%-6.2f|\n", 12.82)
	fmt.Println("----------")

	char := 97
	str := fmt.Sprintf("%c", char)
	fmt.Println(str)

	var message string = "Hello and welcome to"
	var year int = 2022

	var complete_msg = fmt.Sprintf("%s educative in %d", message, year)
	fmt.Println(complete_msg)

}
