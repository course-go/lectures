package main

import (
	"fmt"
	"strconv"
)

// START OMIT

type Slice[T any] []T

func (s Slice[T]) Map[U any](f func(T) U) Slice[U] { // Method with its own type parameter U
	return Map(s, f)
}

func Map[T, U any](s Slice[T], f func(T) U) Slice[U] { // Package scoped function
	result := make(Slice[U], 0, len(s))
	for _, value := range s {
		result = append(result, f(value))
	}
	return result
}

func main() {
	s := Slice[int]{1, 2, 3}
	fmt.Println(Map(s, strconv.Itoa))
	fmt.Println(s.Map(strconv.Itoa))
}

// END OMIT
