### Create an interactive calculator that exhibits the following behavior when it is executed.

Display the following menu
1. Add
2. Subtract
3. Multiply
4. Divide
5. Exit

if user input == 5
   then exit

if user input == 1 to 4
   accept the 2 operands from the user
   perform the corresponding operation
   print the result
   display the menu again

if user input < 1 OR > 5
   print "invalid choice"
   display the menu again

```go
// code from accepting input from the user
	var input int
	fmt.Printf("Enter your choice:")
	fmt.Scanln(&input)

	fmt.Println("Value is: ", input)
```