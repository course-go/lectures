package main

import "fmt"

func main() {
	var array0 [10]int

	array1 := array0

	fmt.Printf("Array 1: %v\n", array0)
	fmt.Printf("Array 2: %v\n", array1)

	for i := 0; i < len(array0); i++ {
		array0[i] = i * 2
	}

	fmt.Println("-------------------------------")
	fmt.Printf("Array 1: %v\n", array0)
	fmt.Printf("Array 2: %v\n", array1)
}
