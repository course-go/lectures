package main

import "fmt"

func main() {
	i := 42
	fmt.Println(i)

	var ptrToInteger *int
	fmt.Println(ptrToInteger)

	ptrToInteger = &i
	fmt.Println(ptrToInteger)
	fmt.Println(*ptrToInteger)

	*ptrToInteger++
	fmt.Println(i)
	fmt.Println(*ptrToInteger)
}
