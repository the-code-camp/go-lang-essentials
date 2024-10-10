package main

import (
	"fmt"
	"math"
)

// public interface Greeting {
// 	Hello()
//}
/**
public class MyHello implements Greeting {
}
*/

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type AreaFinder interface {
	Area() float64
}

type PerimeterFinder interface {
	Perimeter() float64
}

type ShapeFinder interface {
	AreaFinder
	PerimeterFinder
}

func main() {
	c := Circle{10}
	r := Rectangle{10, 12}

	// fmt.Println("Area of Circle: ", c.Area())
	// fmt.Println("Area of Rectangle: ", r.Area())

	printArea(c)
	printArea(r)

	printPerimiter(c)
	printPerimiter(r)

	printShape(c)
	printShape(r)
}

func printShape(x ShapeFinder) {
	fmt.Println("ShapeFinder: ")
	fmt.Println("Area: ", x.Area())
	fmt.Println("Perimeter: ", x.Perimeter())

}

func printPerimiter(x PerimeterFinder) {
	fmt.Println("Perimeter: ", x.Perimeter())
}

func printArea(x AreaFinder) {
	fmt.Println("Area: ", x.Area())
}
