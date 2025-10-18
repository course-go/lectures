package main

import "fmt"

// START OMIT

type Person struct {
	name string
}

func main() {
	p1 := new(Person)
	p1.name = "Bob"
	p2 := &Person{
		name: "Bob",
	}
	fmt.Printf("Person with new: %T %v\n", p1, p1)
	fmt.Printf("Person with ampersand: %T %v\n", p2, p2)
	i1 := new(int)
	// i2 := &0 Invalid
	var integer int
	i2 := &integer
	fmt.Printf("Integer with new: %T %v\n", i1, i1)
	fmt.Printf("Integer with ampersand: %T %v\n", i2, i2)
}

// END OMIT
