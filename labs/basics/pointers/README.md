## Pointers

Go supports pointers, allowing you to pass references to values and records within your program.

```go
package main

import "fmt"

func main() {
	wallet := 1000

	makePurchase(wallet, 200)
	fmt.Println("New wallet value:", wallet)
}

func makePurchase(wallet int, purchaseAmount int) {
	wallet = wallet - purchaseAmount
}
```

By default, arguments are passed by value. `makePurchase` will get a copy of wallet and that will be distinct from the one in the calling function.

This program can be fixed by using pointers:

```go
package main

import "fmt"

func main() {
	wallet := 1000

	makePurchase(&wallet, 200)
	fmt.Println("New wallet value:", wallet)
}

func makePurchase(wallet *int, purchaseAmount int) {
	*wallet = *wallet - purchaseAmount
}
```
After making the changes, `makePurchase` is receiving `wallet` as `pointer`. 

`&wallet` in the `main` function gives the memory address of wallet, i.e. a `pointer to wallet`.