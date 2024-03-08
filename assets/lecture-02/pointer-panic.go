package main

import "fmt"

func main() {
	var ptrToInteger *int
	fmt.Println(ptrToInteger)

	*ptrToInteger++
}
