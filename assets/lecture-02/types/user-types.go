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
	var id ID = 0
	fmt.Println(id)
	n := Name("Jan")
	var s Surname
	s = "Novák"
	fmt.Println(n, s)
}
