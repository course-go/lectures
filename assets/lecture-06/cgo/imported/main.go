package main

// START OMIT

// #include "hello.c"
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello from Go.")
	hello := C.CString("Hello from C.\n")
	err := C.hello_from_c(hello)
	if err == C.eof {
		fmt.Println("Printing in C failed with EOF.")
	}
	C.free(unsafe.Pointer(hello))
}

// END OMIT
