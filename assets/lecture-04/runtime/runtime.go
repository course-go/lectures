package main

import (
	"fmt"
	"runtime"
)

// START OMIT

func main() {
	fmt.Println(`Logical CPUs ("P"s):`, runtime.NumCPU())
	runtime.GC() // Invokes garbage collector
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(8))
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(4))
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	runtime.LockOSThread()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	fmt.Println("Heap:", stats.Alloc, "B")
}

// END OMIT
