package main

import "fmt"

type (
	// ID is user-defined data type
	ID uint32
	// Name is user-defined data type
	Name string
	// Surname is user-defined data type
	Surname string
)

func registerUser(id ID, n Name, s Surname) {
	fmt.Printf("Registering: %d %s %s", id, n, s)
}

func main() {
	var i ID = 1
	var n Name = "Jan"
	var s Surname = "Novák"
	registerUser(i, n, s)
}
