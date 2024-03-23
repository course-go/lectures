package main

import "fmt"

// START OMIT

type unorderable interface {
	int | float64 | string | bool
}

func compare[T unorderable](x, y T) bool {
	return x < y // Bool does not support ordering via '<' operator
}

func main() {
	fmt.Println(compare(true, true))
	fmt.Println(compare(1, 2))
	fmt.Println(compare(1.5, 2.6))
	fmt.Println(compare("Go", "Gophers!"))
}

// END OMIT
