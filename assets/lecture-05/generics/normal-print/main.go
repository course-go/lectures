package main

import "fmt"

// START OMIT

func printValue(value string) {
	fmt.Println(value) // Note that fmt.Println() accepts `any`
}

func main() {
	printValue("course-go.dev")
}

// END OMIT
