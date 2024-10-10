package main

import (
	"fmt"
	"reflect"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	no := 10
	noPtr := &no

	r := &Rectangle{10.0, 15.0}
	fmt.Println("address of rect: ", r)
	fmt.Printf("Address of rect: %p\n", r)
	fmt.Println(reflect.TypeOf(r))
	fmt.Println("address of no: ", noPtr)

	fmt.Println("Area: ", (*r).Area())
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}
