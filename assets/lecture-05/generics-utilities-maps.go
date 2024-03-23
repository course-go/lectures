package main

import "fmt"

// START OMIT

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func Values[M ~map[K]V, K comparable, V any](m M) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}

	return values
}

// MIDDLE OMIT

func main() {
	users := map[int]string{
		1: "Bob",
		2: "Rob",
	}

	fmt.Println(Keys(users))
	fmt.Println(Values(users))

	universities := map[string]bool{
		"FI MUNI":  true,
		"CUNI MFF": false,
	}

	fmt.Println(Keys(universities))
	fmt.Println(Values(universities))
}

// END OMIT
