package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Add")
		fmt.Println("5. Exit")

		var input int
		var num1, num2 float32
		fmt.Printf("Enter your choice:")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Println("Addition")
			fmt.Printf("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Printf("Enter number 2: ")
			fmt.Scanln(&num2)
			fmt.Println("Addition result is: ", num1+num2)
		case 2:
			fmt.Println("Subtract")
			fmt.Printf("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Printf("Enter number 2: ")
			fmt.Scanln(&num2)
			fmt.Println("Subtraction result is: ", num1-num2)
		case 3:
			fmt.Println("Multiply")
			fmt.Printf("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Printf("Enter number 2: ")
			fmt.Scanln(&num2)
			fmt.Println("Multiply result is: ", num1*num2)
		case 4:
			fmt.Println("Division")
			fmt.Printf("Enter number 1: ")
			fmt.Scanln(&num1)
			fmt.Printf("Enter number 2: ")
			fmt.Scanln(&num2)
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
