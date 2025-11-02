package main

// START OMIT

// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	haystack := "The quick brown fox jumps over the lazy dog"
	needle := "quick"
	cHaystack := C.CString(haystack)
	defer C.free(unsafe.Pointer(cHaystack))
	cNeedle := C.CString(needle)
	defer C.free(unsafe.Pointer(cNeedle))
	cFound := C.strstr(cHaystack, cNeedle)
	// defer C.free(unsafe.Pointer(cFound))

	fmt.Println(cHaystack, ":", C.GoString(cHaystack))
	fmt.Println(cNeedle, ":", C.GoString(cNeedle))
	fmt.Println(cFound, ":", C.GoString(cFound))
}

// END OMIT
