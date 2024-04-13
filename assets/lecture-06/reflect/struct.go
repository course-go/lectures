package main

import (
	"fmt"
	"reflect"
)

// START OMIT

func main() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	fmt.Println("t before:", t)

	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	fmt.Println("Fields:")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("\t%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t after:", t)
}

// END OMIT
