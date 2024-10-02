package main

import (
	"fmt"
)

type (
	// ID is user-defined data type
	ID uint32
	// Name is user-defined data type
	Name string
	// Surname is user-defined data type
	Surname string
)

func main() {
	var id ID = 1
	n := Name("Jan")
	var s Surname
	s = "Novák"
	fmt.Println(id, n, s)
	fmt.Printf("%T, %T, %T", id, n, s)
}
