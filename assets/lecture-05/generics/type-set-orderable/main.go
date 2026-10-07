package main

import "fmt"

// START OMIT

type orderable interface {
	int | float64 | string
}

func compare[T orderable](x, y T) bool {
	return x < y // All types in the orderable interface can be ordered via the `<` operator
}

func main() {
	fmt.Println(compare(1, 2))
	fmt.Println(compare(1.5, 2.6))
	fmt.Println(compare("Go", "Gophers!"))
}

// END OMIT
