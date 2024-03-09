package main

import "fmt"

// START OMIT

func main() {
	var slice []int // default value is nil

	for i := 1; i < 11; i++ {
		slice = append(slice, i)
		fmt.Printf("Slice:          %v\n", slice)
		fmt.Printf("Slice length:   %d\n", len(slice))
		fmt.Printf("Slice capacity: %d\n\n", cap(slice))
	}
}

// END OMIT
