package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	ss := s[2:4]

	fmt.Println(s, len(s), cap(s))

	fmt.Println(ss, len(ss), cap(ss))
	ss[0] = 100
	fmt.Println("After change parent: ", s)
	fmt.Println("After change subslice: ", ss)

	ss = append(ss, 200)
	fmt.Println("After append to the subslice what happens in the parent?? ", s)
	fmt.Println("After append subslice: ", ss)

	fmt.Println("USING MAKE")
	s1 := make([]int, 5, 10)
	s1[0] = 10
	s1[1] = 20
	s1 = append(s1, 99, 100, 101, 102, 103, 104)
	fmt.Println(s1, len(s1), cap(s1))
}
