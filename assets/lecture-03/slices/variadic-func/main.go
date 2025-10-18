package main

import "fmt"

// START OMIT

func printArgs(args ...string) {
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}

	fmt.Println()
}

func main() {
	printArgs() // args == nil
	printArgs("A")
	printArgs("H", "J")
	printArgs("X", "Y", "Z")

	args := []string{
		"Hello",
		"World",
		"!",
	}
	printArgs(args...) // spread values
}

// END OMIT
