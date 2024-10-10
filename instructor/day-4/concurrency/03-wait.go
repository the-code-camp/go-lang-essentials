package main

import (
	"fmt"
	"sync"
	"time"
)

// semaphone based counter

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1) // ideally, you should always add 1
	go f1(wg) // this is a request to schedule this as a go routine
	f2()

	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // always defer done in the first line
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 ended")
}

func f2() {
	fmt.Println("f2 invoked")
}
