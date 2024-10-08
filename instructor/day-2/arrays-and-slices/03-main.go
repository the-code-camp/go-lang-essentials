package main

import "fmt"

func main() {
	nosSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("[Slice] Before mutating:", nosSlice)
	multiplyBy2SliceForEachByIndex(nosSlice)
	fmt.Println("[Slice] After mutating:", nosSlice)
}

// this is a common mistake, always use the index of the slice when you want to mutate
// this doesn't work
func multiplyBy2SliceForEach(s []int) {
	for _, n := range s { // it give a copy of the element
		n = n * 2
	}
}

func multiplyBy2SliceForEachByIndex(s []int) {
	for index, _ := range s {
		s[index] = s[index] * 2
	}
}
