package main

import (
	"fmt"
	"math"
)

// Define the Shape interface
type Shape interface {
	Area() float64
}

// Create a Rectangle type that implements the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Create a Circle type that implements the Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Function that takes a Shape as an argument and calculates its area
func CalculateArea(s Shape) {
	area := s.Area()
	fmt.Printf("Area: %f\n", area)
}

func main() {
	rectangle := Rectangle{Width: 4, Height: 3}
	circle := Circle{Radius: 2}

	CalculateArea(rectangle)
	CalculateArea(circle)
}
