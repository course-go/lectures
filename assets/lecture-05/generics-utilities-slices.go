package main

import (
	"fmt"
	"slices"
)

// START OMIT

func MapFunc[T any, U any](slice []T, mapFunc func(T) U) []U {
	mappedSlice := make([]U, 0, len(slice))
	for _, value := range slice {
		mappedSlice = append(mappedSlice, mapFunc(value))
	}

	return mappedSlice
}

func FilterFunc[T any](slice []T, filterFunc func(T) bool) []T {
	filteredSlice := make([]T, 0, len(slice))
	for _, value := range slice {
		if filterFunc(value) {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return slices.Clip(filteredSlice)
}

// MIDDLE OMIT

func main() {
	names := []string{
		"Bob",
		"Rob",
		"X",
		"RoboCop",
	}

	fmt.Println(names)
	names = FilterFunc(names, func(name string) bool {
		return len(name) > 1
	})

	fmt.Println(names)
	names = MapFunc(names, func(name string) string {
		if len(name) > 4 {
			return name[:4]
		}
		return name
	})

	fmt.Println(names)
}

// END OMIT
