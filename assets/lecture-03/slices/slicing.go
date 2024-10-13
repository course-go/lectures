package main

import "fmt"

// START OMIT

func main() {
	var array [10]int
	slice0 := array[4:9]
	slice1 := slice0[3:] // slices can also be sliced

	fmt.Printf("Array:            %v\n", array)
	fmt.Printf("Array length:     %d\n\n", len(array))

	fmt.Printf("Slice 1:          %v\n", slice0)
	fmt.Printf("Slice 1 length:   %d\n", len(slice0))
	fmt.Printf("Slice 1 capacity: %d\n\n", cap(slice0))

	fmt.Printf("Slice 2:          %v\n", slice1)
	fmt.Printf("Slice 2 length:   %d\n", len(slice1))
	fmt.Printf("Slice 2 capacity: %d\n\n", cap(slice1))

	slice1[0] = 99
	slice1[1] = 99

	fmt.Printf("Array:            %v\n", array)
	fmt.Printf("Slice 1:          %v\n", slice0)
	fmt.Printf("Slice 2:          %v\n", slice1)
}

// END OMIT
