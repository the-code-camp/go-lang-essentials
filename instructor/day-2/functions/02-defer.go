package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	// Defer the closing of the file
	defer file.Close()

	b := make([]byte, 100)
	var n int
	if n, err = file.Read(b); err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	fmt.Println("Contents: ", string(b))
	fmt.Println("Number of bytes read: ", n)
}
