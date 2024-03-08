package main

import "fmt"

func execute(function func()) {
	function()
}

func printingFunc() func() {
	return func() {
		fmt.Println("Hello! I'm a printing function!")
	}
}

func main() {
	function := printingFunc()
	execute(function)
}
