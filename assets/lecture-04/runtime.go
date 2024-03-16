package main

import (
	"fmt"
	"runtime"
)

// START OMIT

func main() {
	fmt.Printf("Logical CPUs (\"P\"s): %d\n", runtime.NumCPU())
	runtime.GC() // Invokes garbage collector
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(8))
}

// END OMIT
