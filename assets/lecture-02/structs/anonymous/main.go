package main

import "fmt"

func main() {
	anonymousUser := struct {
		id      int
		name    string
		surname string
	}{
		1,
		"Linus",
		"Torvalds",
	}
	fmt.Println(anonymousUser)
}
