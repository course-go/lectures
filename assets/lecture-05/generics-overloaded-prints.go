package main

import "fmt"

// START OMIT

func printValue(value string) {
	fmt.Println(value)
}

func printValue(value int) {
	fmt.Println(value)
}

func main() {
	printValue("We need more Gophers!")
	printValue(42)
}

// END OMIT
