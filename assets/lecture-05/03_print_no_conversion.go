package main

import "fmt"

type Value string

func printValue(value Value) {
	fmt.Println(value)
}

func main() {
	v := "Go language!" // string
	printValue(v)
}
