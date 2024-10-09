package main

import (
	"fmt"

	"github.com/the-code-camp/go-lang-essentials/utils"
	q "rsc.io/quote"
)

func main() {
	result := utils.Sum(10, 20)
	fmt.Println(q.Glass())
	fmt.Println("Result is: ", result)
}
