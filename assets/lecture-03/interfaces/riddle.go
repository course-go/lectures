package main

import "fmt"

// START OMIT

// I represents a new interface type
type I interface{}

// T represents an user-defined data type
type T struct{}

func main() {
	var t *T
	if t == nil {
		fmt.Println("t is nil")
	} else {
		fmt.Println("t is not nil")
	}

	var i I = t
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is not nil")
	}
}

// END OMIT
