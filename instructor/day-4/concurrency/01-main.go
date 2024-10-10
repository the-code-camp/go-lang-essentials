package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // this is a request to schedule this as a go routine
	f2()

	// time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 ended")
}

func f2() {
	fmt.Println("f2 invoked")
}
