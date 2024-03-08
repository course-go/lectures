package main

import "fmt"

func swap(a, b int) (x, y int) {
	y = a
	x = b
	return
}

func main() {
	a, b := 3, 4
	fmt.Println("a: ", a, ", b: ", b)
	a, b = swap(a, b)
	fmt.Println("a: ", a, ", b: ", b)
}
