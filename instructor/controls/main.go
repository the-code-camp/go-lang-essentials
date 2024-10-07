package main

import (
	"fmt"
	"runtime"
)

func main() {
	// labelsAndGoTo()

	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func labelsAndGoTo() {
	i := 1
	for {
		if i > 10 {
			goto EXIT
		}
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}
EXIT: //labels
	fmt.Println("program finished...")

}
