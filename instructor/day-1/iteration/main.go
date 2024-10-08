package main

import "fmt"

func main() {
	i := 0
	for ; i <= 10; i++ {
		fmt.Println(i)
	}

	for i > 0 {
		fmt.Println(i)
		i--
	}

	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c ", ch)
	}

	fmt.Println()
	for rows := 1; rows <= 5; rows++ {
		for cols := 1; cols <= rows; cols++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
