package main

import "fmt"

// START OMIT

type Value string

func printValue(value Value) {
	fmt.Println(value)
}

func main() {
	v := "We need more Gophers!" // string
	printValue(v)
}

// END OMIT
