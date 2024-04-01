package main

import "fmt"

// START OMIT

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

// MIDDLE OMIT

func main() {
	var u User
	u.id = 1
	u.name = "Linus"
	u.surname = "Torvalds"
	fmt.Println(u)

	var pUser *User
	pUser = &u
	fmt.Println(pUser)
	fmt.Println(*pUser)

	(*pUser).id = 10000
	fmt.Println(*pUser)

	pUser.id = 20000
	fmt.Println(*pUser)
}

// END OMIT
