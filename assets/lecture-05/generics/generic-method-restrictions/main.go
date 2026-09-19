package main

import "fmt"

// START OMIT

type Slice[T any] []T

func (s Slice[T]) Map[U any](f func(T) U) Slice[U] {
	result := make(Slice[U], 0, len(s))
	for _, value := range s {
		result = append(result, f(value))
	}
	return result
}

type Mapper interface {
	Map[U any](f func(int) U) Slice[U] // Interface methods may not declare type parameters
}

func main() {
	var m Mapper = Slice[int]{1, 2, 3} // A generic method cannot implement an interface method
	fmt.Println(m)
}

// END OMIT
