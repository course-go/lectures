package main

import "fmt"

func main() {
	if j := 12; j > 16 {
		fmt.Println(j, "> 16")
	} else if j < 8 {
		fmt.Println(j, "< 8")
	} else {
		fmt.Println("8 <", j, "< 16")
	}
	// "j" is no longer in scope
}
