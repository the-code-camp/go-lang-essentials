package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// add(100, 200)
	// subtract(100, 200)

	// log.Println("operation started")
	// add(100, 200)
	// log.Println("operation completed")

	// log.Println("operation started")
	// subtract(100, 200)
	// log.Println("operation completed")

	// logAdd(100, 200)
	// logSubtract(100, 200)

	// DRY -> Dont repeat yourself -> removing duplicacy

	// logOperation(add, 100, 200)
	// logOperation(subtract, 100, 200)

	// logAdd := getLogOperation(add)
	// logAdd(100, 200)

	// logSubtract := getLogOperation(subtract)
	// logSubtract(100, 200)

	// we want to profile operations
	// profileAdd := getProfileOperation(add)
	// profileAdd(100, 200)

	// profileSub := getProfileOperation(subtract)
	// profileSub(100, 200)

	logAdd := getLogOperation(add)
	profileLogAdd := getProfileOperation(logAdd)
	profileLogAdd(100, 200)

	fmt.Println()

	profileAdd := getProfileOperation(add)
	logProfileAdd := getLogOperation(profileAdd)
	logProfileAdd(100, 200)
}

func getProfileOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed: ", elapsed)
	}
}

func getLogOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("operation started")
		operation(x, y)
		log.Println("operation completed")
	}
}

func logOperation(operation func(int, int), x, y int) {
	log.Println("operation started")
	operation(x, y)
	log.Println("operation completed")
}

// func logAdd(x, y int) {
// 	log.Println("operation started")
// 	add(x, y)
// 	log.Println("operation completed")
// }

// func logSubtract(x, y int) {
// 	log.Println("operation started")
// 	subtract(x, y)
// 	log.Println("operation completed")
// }

func add(x, y int) {
	fmt.Println("Add result is: ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result is: ", x-y)
}
