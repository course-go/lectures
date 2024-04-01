package main

import "fmt"

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

func printUser(u User) {
	fmt.Println(u)
}

func main() {
	user1 := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
	}
	printUser(user1)
}
