package main

import (
	"fmt"
	"math"
)

// START OMIT

type ClosedShape interface {
	area() float64
}

func area(shape ClosedShape) float64 {
	return shape.area()
}

type Circle struct {
	x, y   float64
	radius float64
}

type Rectangle struct {
	x, y          float64
	width, height float64
}

// MIDDLE OMIT

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	fmt.Println("Rectangle")
	r := Rectangle{x: 0, y: 0, width: 100, height: 100}
	fmt.Println(r)
	fmt.Println(area(r))
	fmt.Println(r.area())
	fmt.Println()

	fmt.Println("Circle")
	c := Circle{x: 0, y: 0, radius: 100}
	fmt.Println(c)
	fmt.Println(area(c))
	fmt.Println(c.area())
	fmt.Println()
}

// END OMIT
