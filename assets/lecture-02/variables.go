package main

import "fmt"

var (
	a int    = 3
	x string = "Hello world"
)

func main() {
	var b int = 4
	c := 5
	c = 6
	fmt.Println("a:", a, "b:", b, "c:", c)
}
