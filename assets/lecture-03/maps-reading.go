package main

import "fmt"

// START OMIT

func main() {
	m1 := make(map[string]string)
	fmt.Println(m1)
	m1["one"] = "I"
	m1["two"] = "II"
	m1["three"] = "III"
	m1["four"] = "IV"
	m1["five"] = "V"

	value, exists := m1["two"] // comma ok idiom
	if exists {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	value, exists = m1["thousand"] // comma ok idiom
	if exists {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}
}

// END OMIT
