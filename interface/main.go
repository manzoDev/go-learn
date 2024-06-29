package main

import "fmt"


type Shape interface {
	Area() float64
	Perimeter() float64
}


func main() {

	c := Circle{Radius: 5}

	fmt.Println("Original Circle:")
	PrintShapeProperties(c)

	r := Rectangle{Width: 2, Height: 2}

	fmt.Println("Original Rectangle:")
	PrintShapeProperties(r)
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 3.14 * r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return  2 * (r.Width + r.Height) 
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return  2 * 3.14 * c.Radius 
}

func PrintShapeProperties(s Shape){ 
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}