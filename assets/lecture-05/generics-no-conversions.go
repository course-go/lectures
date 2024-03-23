package main

import "fmt"

// START OMIT

type numeric interface {
	int | float64 | complex128
}

func add[T numeric](x, y T) T {
	return x + y
}

type myInt int // Alias for `int`

func main() {
	x := myInt(42)
	fmt.Println(add(x, x)) // Underlying type of `myInt` is `int`, shouldn't it work?
}

// END OMIT
