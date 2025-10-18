package main

import "fmt"

// START OMIT

func main() {
	var array [100]byte
	fmt.Printf("Array length:   %d\n\n", len(array))

	slice0 := array[:]
	fmt.Printf("Slice 0 length:   %d\n", len(slice0))
	fmt.Printf("Slice 0 capacity: %d\n\n", cap(slice0))

	slice1 := array[10:20]
	fmt.Printf("Slice 1 length:   %d\n", len(slice1))
	fmt.Printf("Slice 1 capacity: %d\n\n", cap(slice1))

	slice2 := array[60:]
	fmt.Printf("Slice 2 length:   %d\n", len(slice2))
	fmt.Printf("Slice 2 capacity: %d\n\n", cap(slice2))
}

// END OMIT
