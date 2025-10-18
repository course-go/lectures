package main

import "fmt"

// START OMIT

type Slice[T any] []T // Generic structure

func (s Slice[T]) Print() { // Method
	for _, value := range s {
		fmt.Println(value)
	}
}

func Print[T any](s Slice[T]) { // Function
	for _, value := range s {
		fmt.Println(value)
	}
}

func main() {
	s := Slice[int]{1, 2, 3}
	Print(s)
	fmt.Println()
	s.Print()
}

// END OMIT
