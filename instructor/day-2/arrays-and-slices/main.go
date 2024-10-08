package main

import "fmt"

func main() {
	nosArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("[Array] Before mutating:", nosArray)
	multiplyBy2(nosArray)
	fmt.Println("[Array] After mutating:", nosArray)

	nosSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("[Slice] Before mutating:", nosSlice)
	multiplyBy2Slice(nosSlice)
	fmt.Println("[Slice] After mutating:", nosSlice)
}

// Go is pass by value

func multiplyBy2(s [5]int) {
	for i := 0; i < 5; i++ {
		s[i] = s[i] * 2
	}
}

func multiplyBy2Slice(s []int) {
	for i := 0; i < 5; i++ {
		s[i] = s[i] * 2
	}
}
