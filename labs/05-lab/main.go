package main

import (
	"fmt"
	"time"
)

/*
Modify the program to print the values in order of receiving

  - 500 (this is produced in 2 seconds)

  - 200 (this is produced in 5 seconds)

    This should be done using the language constructs we have learned so far. This
    can also be done using select case but we dont have to use that.
*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 200
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 500
	}()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
}
