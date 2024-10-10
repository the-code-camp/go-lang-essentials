package main

import (
	"fmt"
	"sync"
	"time"
)

// semaphone based counter
var wg sync.WaitGroup

func main() {
	wg.Add(1) // ideally, you should always add 1
	go f1()   // this is a request to schedule this as a go routine
	f2()

	wg.Wait()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 ended")
	wg.Done() // it will always be reduced by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
