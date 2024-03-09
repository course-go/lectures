package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[0] = ""
	m[1] = "1st"
	m[2] = "2nd"
	m[3] = "3rd"

	fmt.Printf("'%s'\n", m[0])
	fmt.Printf("'%s'\n", m[1])
	fmt.Printf("'%s'\n", m[2])
	fmt.Printf("'%s'\n", m[3])
	fmt.Printf("'%s'\n", m[4])
}
