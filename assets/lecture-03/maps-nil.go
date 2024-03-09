package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1)

	fmt.Println(m1["one"])

	m1["zero"] = 0
}
