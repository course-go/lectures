package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[0] = "one"
	m[1] = "two"
	m[2] = "three"
	m[3] = "four"

	for key, val := range m {
		fmt.Println(key, val)
	}

	fmt.Println()

	fmt.Println("Length:", len(m))
}
