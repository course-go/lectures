package main

import "fmt"

// START OMIT

func printValue[T any](value T) {
	fmt.Println(value)
}

func main() {
	printValue[int]("course-go.dev")
	printValue[[]string]('*')
	printValue[string](42)
	printValue[int](3.14)
	printValue[byte](1 + 2i)
	printValue[[]byte]([]int{1, 2, 3})
}

// END OMIT
