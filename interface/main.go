package main

import "fmt"


func main() {

	c := Circle{Radius: 5}

	fmt.Println("Original Circle:")
	PrintShapeProperties(c)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return  2 * 3.14 * c.Radius 
}

func PrintShapeProperties(c Circle){ 
	fmt.Printf("Area: %f\n", c.Area())
	fmt.Printf("Perimeter: %f\n", c.Perimeter())
}