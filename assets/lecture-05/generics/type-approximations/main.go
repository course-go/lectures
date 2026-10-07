package main

import "fmt"

// START OMIT

type numeric interface {
	~int | ~float64 | ~complex128
}

func add[T numeric](x, y T) T {
	return x + y
}

type myInt int // Defined type with `int` as its underlying type (not an alias)

func main() {
	x := myInt(42)
	fmt.Println(add(x, x))
}

// END OMIT
