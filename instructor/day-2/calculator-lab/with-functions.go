package main

import "fmt"

func getInputFor(operation string) (num1, num2 float32) {
	fmt.Println(operation)
	fmt.Printf("Enter number 1: ")
	fmt.Scanln(&num1)
	fmt.Printf("Enter number 2: ")
	fmt.Scanln(&num2)
	return num1, num2
}

func printMenuAndHandleInput() int {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Add")
	fmt.Println("5. Exit")

	var input int
	fmt.Printf("Enter your choice:")
	fmt.Scanln(&input)
	return input
}

func main() {
	for {
		input := printMenuAndHandleInput()
		switch input {
		case 1:
			num1, num2 := getInputFor("Addition")
			fmt.Println("Addition result is: ", num1+num2)
		case 2:
			num1, num2 := getInputFor("Subtration")
			fmt.Println("Subtraction result is: ", num1-num2)
		case 3:
			num1, num2 := getInputFor("Multiply")
			fmt.Println("Multiply result is: ", num1*num2)
		case 4:
			num1, num2 := getInputFor("Division")
			fmt.Printf("Division result is: %.2f\n", num1/num2)
		case 5:
			goto EXIT
		default:
			fmt.Println("Invalid choice..., retry")
		}
	}
EXIT:
	fmt.Println("Program finished")
}
