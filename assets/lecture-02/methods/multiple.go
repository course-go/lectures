package main

import "fmt"

// START OMIT

// Line represents a line in 2D plane.
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func (l *Line) translate(dx, dy float64) {
	fmt.Printf("Translating line %v by %f %f\n", *l, dx, dy)
}

func (l *Line) rotate(angle float64) {
	fmt.Printf("Rotating line %v by %f\n", *l, angle)
}

func main() {
	l := Line{x1: 0, y1: 0, x2: 100, y2: 100}
	l.translate(5, 5)
	l.rotate(90)
}

// END OMIT
