package main

import "fmt"

// START OMIT

func printString(value string) {
	fmt.Println(value)
}

func printInt(value int) {
	fmt.Println(value)
}

func main() {
	printString("We need more Gophers!")
	printInt(60045382)
}

// END OMIT
