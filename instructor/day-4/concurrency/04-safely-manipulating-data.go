package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

// var count int
var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count: ", counter.count)
}

var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
	// mutex.Lock()
	// {
	// 	count++
	// }
	// mutex.Unlock()
}
