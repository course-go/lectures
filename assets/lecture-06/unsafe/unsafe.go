package main

import (
	"fmt"
	"unsafe"
)

// START OMIT

func main() {
	type s struct {
		a string
		b int32
		c uint64
	}
	x := s{a: "Hello", b: -72, c: 1000}
	fmt.Println(unsafe.Sizeof(x))
	fmt.Println(unsafe.Alignof(x))
	fmt.Println(unsafe.Offsetof(x.a))

	pointer := unsafe.Pointer(&x)
	a := (*string)(unsafe.Add(pointer, unsafe.Offsetof(x.a)))
	b := (*int32)(unsafe.Add(pointer, unsafe.Offsetof(x.b)))
	fmt.Println("a:", *a, "b", *b)

	truth := []string{"Go", "is", "the", "best!"}
	fmt.Println(unsafe.Sizeof(truth))
	slicePointer := unsafe.Pointer(unsafe.SliceData(truth))
	itemSize := unsafe.Sizeof("")
	for i := range len(truth) {
		fmt.Println(*(*string)(unsafe.Add(slicePointer, uintptr(i)*itemSize)))
	}
}

// END OMIT
