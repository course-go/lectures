package main

import (
	"fmt"
	"reflect"
)

// START OMIT

func main() {
	var x rune = 'x'
	vx := reflect.ValueOf(x)
	fmt.Println("type:", vx.Type())
	fmt.Println("kind is int32:", vx.Kind() == reflect.Int32)
	x = rune(vx.Int())
	fmt.Println("x:", x)

	fmt.Println()

	var y float64 = 1.7
	p := reflect.ValueOf(&y)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	vy := p.Elem()
	fmt.Println("settability of v:", vy.CanSet())
	vy.SetFloat(7.1)
	fmt.Println(vy.Interface())
	fmt.Println(y)
}

// END OMIT
