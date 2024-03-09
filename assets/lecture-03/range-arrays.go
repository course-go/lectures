package main

import "fmt"

// START OMIT

func main() {
	array := [4]string{"one", "two", "three", "four"}

	for index := range array {
		fmt.Println(index)
	}

	fmt.Println()

	for index, item := range array {
		fmt.Println(index, item)
	}

	fmt.Println()

	for _, item := range array {
		fmt.Println(item)
	}
}

// END OMIT
