package main

import "fmt"

func main() {
	var array0 [10]byte
	array1 := [10]int{1, 2, 3}
	array2 := [10]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	fmt.Printf("Array 1:\n  length: %d\n  content: %v\n", len(array0), array0)
	fmt.Printf("Array 2:\n  length: %d\n  content: %v\n", len(array1), array1)
	fmt.Printf("Array 3:\n  length: %d\n  content: %v\n", len(array2), array2)

	var array [10]int
	fmt.Printf("Original array: %v\n", array)

	for i := 0; i < len(array0); i++ {
		array[i] = i * 2
	}

	fmt.Printf("Modified array: %v\n", array)
}
