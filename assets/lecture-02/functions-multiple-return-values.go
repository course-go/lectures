package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	fmt.Println("a: ", a, "b: ", b)
	a, b = swap(a, b)
	fmt.Println("a: ", a, "b: ", b)
}
