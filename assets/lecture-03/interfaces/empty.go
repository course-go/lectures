package main

import "fmt"

// I represents a new interface type (in this case empty interface)
type I interface{}

// T represents an user-defined data type
type T struct{}

func main() {
	var i I = T{}
	fmt.Printf("Value: %v, Type: %T", i, i)
}
