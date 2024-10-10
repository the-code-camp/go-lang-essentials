package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	count := flag.Int("count", 0, "Number of goroutines to spin up")
	flag.Parse()

	fmt.Printf("Starting %d go routines. Press ENTER \n", *count)
	fmt.Scanln()

	wg := &sync.WaitGroup{}
	for i := 1; i <= *count; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Printf("Press ENTER to stop\n", *count)
	fmt.Scanln()

}

func fn(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started\n", i)
	// r := rand.Intn(10)
	time.Sleep(3 * time.Second)
	fmt.Printf("fn[%d] ended\n", i)
}
