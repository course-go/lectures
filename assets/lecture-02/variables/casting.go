package main

import "fmt"

func main() {
	var a float64
	a = float64(5)
	p := (*int)(nil)

	fmt.Printf("a: %f, p: %p", a, p)

	x := byte(24)
	y := int8(12)
	z := 12 // int

	fmt.Printf("sum: %d\n", x+byte(y)+byte(z))
}
