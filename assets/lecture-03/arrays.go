package main

import "fmt"

func main() {
	var a1 [10]byte
	a2 := [10]int{1, 2, 3}
	a3 := [10]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	fmt.Printf("Array 1:\n  length: %d\n  content: %v\n", len(a1), a1)
	fmt.Printf("Array 2:\n  length: %d\n  content: %v\n", len(a2), a2)
	fmt.Printf("Array 3:\n  length: %d\n  content: %v\n", len(a3), a3)

	var a [10]int
	fmt.Printf("Original array: %v\n", a)

	for i := 0; i < len(a1); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Modified array: %v\n", a)
}
