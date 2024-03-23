package main

import "fmt"

// START OMIT

func printValue(value interface{}) {
	fmt.Println(value) // No type switching here as fmt.Println() accepts `any`
}

func main() {
	printValue("course-go.dev")
	printValue('*')
	printValue(42)
	printValue(3.14)
	printValue(1 + 2i)
	printValue([]int{1, 2, 3})
}

// END OMIT
