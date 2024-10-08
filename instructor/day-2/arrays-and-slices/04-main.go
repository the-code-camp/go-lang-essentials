package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = slices.Delete(s, 2, 4)
	fmt.Println(s)

	// more performant, but doesn't maintain the order
	s1 := []int{1, 2, 3, 4, 5}
	s1[2] = s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Println(s1)
}
