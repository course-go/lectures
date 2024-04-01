package main

import "fmt"

// START OMIT

func main() {
	a := make([]int, 10, 16)
	b := make([]int, 8)

	fmt.Printf("Slice:          %v\n", a)
	fmt.Printf("Slice length:   %d\n", len(a))
	fmt.Printf("Slice capacity: %d\n\n", cap(a))

	fmt.Printf("Slice:          %v\n", b)
	fmt.Printf("Slice length:   %d\n", len(b))
	fmt.Printf("Slice capacity: %d\n\n", cap(b))
}

// END OMIT
