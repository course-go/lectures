package main

import "fmt"

// START OMIT

func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	printValue[string]("course-go.dev")
	printValue[rune]('*')
	printValue[int](42)
	printValue[float32](3.14)
	printValue[complex64](1 + 2i)
	printValue[[]int]([]int{1, 2, 3})
}

// END OMIT
