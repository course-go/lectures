package main

import (
	"fmt"
	"time"
)

func print(s string) {
	for range 5 {
		fmt.Printf("Hello from %s!\n", s)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go print("first")
	go print("second")
	print("main")
}
