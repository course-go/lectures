package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	m2 := make(map[int]string, 256)
	fmt.Println(m1)
	fmt.Println(m2)

	m1[1] = "one"
	m1[2] = "two"
	m1[3] = "three"
	m1[4] = "four"

	fmt.Println(m1)
}
