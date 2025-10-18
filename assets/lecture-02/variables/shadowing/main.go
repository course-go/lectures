package main

import "fmt"

func main() {
	x := "world"
	{
		x := "hello"
		fmt.Println(x)
	}

	fmt.Println(x)
}
