package main

import "fmt"

// START OMIT

func init() {
	fmt.Println("Init gets called first!")
}

func main() {
	fmt.Println("Followed by the main function!")
}

// END OMIT
