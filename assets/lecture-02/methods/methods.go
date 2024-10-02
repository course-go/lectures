package main

import (
	"fmt"
	"math"
)

// START OMIT

// Line represents a line in 2D plane.
type Line struct {
	x1, y1 float64
	x2, y2 float64
}

func (l Line) length() float64 {
	return math.Hypot(l.x1-l.x2, l.y1-l.y2)
}

func main() {
	l := Line{x1: 0, y1: 0, x2: 100, y2: 100}
	fmt.Println(l)
	length := l.length()
	fmt.Println(length)
}

// END OMIT
