package main

import "fmt"

func main() {
	for i := range 5 {
		fmt.Println(i)
	}

	fmt.Println()

	for range 3 {
		fmt.Println("X")
	}
}
