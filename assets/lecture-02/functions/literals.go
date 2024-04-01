package main

import "fmt"

func main() {
	decrement := func(a, b int) int {
		return a - b
	}(10, 7) // directly invoked
	fmt.Println(decrement)

	sum := func(a, b int) int {
		return a + b
	} // assigned to variable
	fmt.Println(sum(3, 4))
}
