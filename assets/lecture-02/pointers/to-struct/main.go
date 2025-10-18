package main

import "fmt"

// User is a new data type
type User struct {
	id      uint32
	name    string
	surname string
}

// START OMIT

func main() {
	u := User{
		id:      1,
		name:    "Linus",
		surname: "Torvalds",
	}
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
