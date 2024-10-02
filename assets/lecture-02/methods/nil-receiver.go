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

func (l *Line) length() float64 {
	if l == nil {
		return 0
	}

	return math.Hypot(l.x1-l.x2, l.y1-l.y2)
}

func main() {
	var l *Line
	fmt.Printf("(%v, %T)\n", l, l)
	fmt.Println(l.length())
}

// END OMIT
