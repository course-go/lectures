package main

import "fmt"

func swap(a int, b int) (x int, y int) {
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
