package main

import "fmt"

func main() {
	numbers := [5]int{10, 12, 13, 14, 15}
	fmt.Println("Length of array:", len(numbers))
	for _, n := range numbers {
		fmt.Println(n)
	}

	var s []int
	s = append(s, 14, 15, 16)
	fmt.Println("Length of slice: ", len(s), s)

	s2 := make([]int, 5)
	s2 = append(s2, 10)
	s2[0] = 20
	s2[4] = 23
	fmt.Println("Length of slice: ", len(s2), s2)

	fmt.Println("value ", s2[4])

	// subslicing
	brothers := []string{"chico", "harpo", "groucho", "gummo", "zeppo"}
	fmt.Println("All brothers: ", brothers)
	firstThree := brothers[0:3]
	fmt.Println("First Three: ", firstThree)

	lastFour := brothers[1:]
	fmt.Println("Last Four: ", lastFour)

	firstFour := brothers[:4]
	fmt.Println("First Four: ", firstFour)

}
