package main

import "fmt"

// Adder is an interface with one method signature
type Adder interface {
	add(int, int) int
}

// Calculator is a user-defined data types that satisfy Adder interface
type Calculator struct {
}

func (c Calculator) add(x, y int) int {
	return x + y
}

func main() {
	a := Adder(Calculator{})
	fmt.Printf("Result is %d\n", a.add(1, 2))
}
