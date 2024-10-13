package main

import "fmt"

// START OMIT

func main() {
	var array [10]int
	slice := array[:]

	fmt.Printf("Array before modification: %v\n", array)
	fmt.Printf("Slice before modification: %v\n\n", slice)

	for i := 0; i < len(array); i++ {
		array[i] = i * 2 // modifications via array
	}

	fmt.Printf("Array after modification:  %v\n", array)
	fmt.Printf("Slice after modification:  %v\n\n", slice)

	for i := 0; i < len(slice); i++ {
		slice[i] = 42 // modifications via slice
	}

	fmt.Printf("Array after modification:  %v\n", array)
	fmt.Printf("Slice after modification:  %v\n", slice)
}

// END OMIT
