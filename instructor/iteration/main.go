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
}
